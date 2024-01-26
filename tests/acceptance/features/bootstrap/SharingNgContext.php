<?php declare(strict_types=1);
/**
 * @author Viktor Scharf <scharf.vi@gmail.com>
 *
 * @copyright Copyright (c) 2023, ownCloud GmbH
 * @license AGPL-3.0
 *
 * This code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, version 3,
 * as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License, version 3,
 * along with this program.  If not, see <http://www.gnu.org/licenses/>
 *
 */

use Behat\Behat\Context\Context;
use Behat\Behat\Hook\Scope\BeforeScenarioScope;
use Psr\Http\Message\ResponseInterface;
use TestHelpers\GraphHelper;
use TestHelpers\WebDavHelper;
use Behat\Gherkin\Node\TableNode;
use PHPUnit\Framework\Assert;

require_once 'bootstrap.php';

/**
 * Acceptance test steps related to testing sharing ng features
 */
class SharingNgContext implements Context {
	private FeatureContext $featureContext;
	private SpacesContext $spacesContext;

	/**
	 * This will run before EVERY scenario.
	 * It will set the properties for this object.
	 *
	 * @BeforeScenario
	 *
	 * @param BeforeScenarioScope $scope
	 *
	 * @return void
	 */
	public function before(BeforeScenarioScope $scope): void {
		// Get the environment
		$environment = $scope->getEnvironment();
		// Get all the contexts you need in this context from here
		$this->featureContext = $environment->getContext('FeatureContext');
		$this->spacesContext = $environment->getContext('SpacesContext');
	}

	/**
	 * @param string $user
	 * @param TableNode|null $body
	 *
	 * @return ResponseInterface
	 * @throws Exception
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function createLinkShare(string $user, TableNode $body): ResponseInterface {
		$bodyRows = $body->getRowsHash();
		$space = $bodyRows['space'];
		$resourceType = $bodyRows['resourceType'];
		$resource = $bodyRows['resource'];

		$spaceId = ($this->spacesContext->getSpaceByName($user, $space))["id"];
		if ($resourceType === 'folder') {
			$itemId = $this->spacesContext->getResourceId($user, $space, $resource);
		} else {
			$itemId = $this->spacesContext->getFileId($user, $space, $resource);
		}

		$bodyRows['displayName'] = $bodyRows['displayName'] ?? null;
		$bodyRows['expirationDateTime'] = $bodyRows['expirationDateTime'] ?? null;
		$bodyRows['password'] = $bodyRows['password'] ?? null;
		$body = [
			'type' => $bodyRows['permissionsRole'],
			'displayName' => $bodyRows['displayName'],
			'expirationDateTime' => $bodyRows['expirationDateTime'],
			'password' => $this->featureContext->getActualPassword($bodyRows['password'])
		];

		return GraphHelper::createLinkShare(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getStepLineRef(),
			$user,
			$this->featureContext->getPasswordForUser($user),
			$spaceId,
			$itemId,
			\json_encode($body)
		);
	}

	/**
	 * @When /^user "([^"]*)" gets permissions list for (folder|file) "([^"]*)" of the space "([^"]*)" using the Graph API$/
	 *
	 * @param string $user
	 * @param string $fileOrFolder   (file|folder)
	 * @param string $resource
	 * @param string $space
	 *
	 * @return void
	 * @throws Exception
	 */
	public function theUserPermissionsListOfResource(string $user, string $fileOrFolder, string $resource, string $space):void {
		$spaceId = ($this->spacesContext->getSpaceByName($user, $space))["id"];

		if ($fileOrFolder === 'folder') {
			$itemId = $this->spacesContext->getResourceId($user, $space, $resource);
		} else {
			$itemId = $this->spacesContext->getFileId($user, $space, $resource);
		}
		$this->featureContext->setResponse(
			GraphHelper::getPermissionsList(
				$this->featureContext->getBaseUrl(),
				$this->featureContext->getStepLineRef(),
				$user,
				$this->featureContext->getPasswordForUser($user),
				$spaceId,
				$itemId
			)
		);
	}

	/**
	 * @param string $user
	 * @param TableNode $table
	 *
	 * @return ResponseInterface
	 *
	 * @throws JsonException
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 * @throws Exception
	 */
	public function sendShareInvitation(string $user, TableNode $table): ResponseInterface {
		$rows = $table->getRowsHash();
		$spaceId = ($this->spacesContext->getSpaceByName($user, $rows['space']))["id"];

		// for resharing a resource, "item-id" in API endpoint takes shareMountId
		if ($rows['space'] === 'Shares') {
			$itemId = GraphHelper::getShareMountId(
				$this->featureContext->getBaseUrl(),
				$this->featureContext->getStepLineRef(),
				$user,
				$this->featureContext->getPasswordForUser($user),
				$rows['resource']
			);
		} else {
			$itemId = ($rows['resourceType'] === 'folder')
				? $this->spacesContext->getResourceId($user, $rows['space'], $rows['resource'])
				: $this->spacesContext->getFileId($user, $rows['space'], $rows['resource']);
		}

		if (\array_key_exists('shareeId', $rows)) {
			$shareeIds[] = $rows['shareeId'];
			$shareTypes[] = $rows['shareType'];
		} else {
			$sharees = array_map('trim', explode(',', $rows['sharee']));
			$shareTypes = array_map('trim', explode(',', $rows['shareType']));

			foreach ($sharees as $sharee) {
				// for non-exiting group or user, generate random id
				$shareeIds[] = $this->featureContext->getAttributeOfCreatedUser($sharee, 'id')
					?: ($this->featureContext->getAttributeOfCreatedGroup($sharee, 'id') ?: WebDavHelper::generateUUIDv4());
			}
		}

		$permissionsRole = $rows['permissionsRole'] ?? null;
		$permissionsAction = $rows['permissionsAction'] ?? null;
		$expireDate = $rows["expireDate"] ?? null;

		$response = GraphHelper::sendSharingInvitation(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getStepLineRef(),
			$user,
			$this->featureContext->getPasswordForUser($user),
			$spaceId,
			$itemId,
			$shareeIds,
			$shareTypes,
			$permissionsRole,
			$permissionsAction,
			$expireDate
		);
		if ($response->getStatusCode() === 200) {
			$this->featureContext->shareNgAddToCreatedUserGroupShares($response);
		}
		return $response;
	}

	/**
	 * @Given /^user "([^"]*)" has sent the following share invitation:$/
	 *
	 * @param string $user
	 * @param TableNode $table
	 *
	 * @return void
	 * @throws Exception
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userHasSentTheFollowingShareInvitation(string $user, TableNode $table): void {
		$response = $this->sendShareInvitation($user, $table);
		$this->featureContext->theHTTPStatusCodeShouldBe(200, "", $response);
	}

	/**
	 * @When /^user "([^"]*)" sends the following share invitation using the Graph API:$/
	 * @When /^user "([^"]*)" tries to send the following share invitation using the Graph API:$/
	 *
	 * @param string $user
	 * @param TableNode $table
	 *
	 * @return void
	 * @throws Exception
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userSendsTheFollowingShareInvitationUsingTheGraphApi(string $user, TableNode $table): void {
		$this->featureContext->setResponse(
			$this->sendShareInvitation($user, $table)
		);
	}

	/**
	 * @When /^user "([^"]*)" creates the following link share using the Graph API:$/
	 *
	 * @param string $user
	 * @param TableNode|null $body
	 *
	 * @return void
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userCreatesAPublicLinkShareWithSettings(string $user, TableNode  $body):void {
		$response = $this->createLinkShare($user, $body);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @Given /^user "([^"]*)" has created the following link share:$/
	 *
	 * @param string $user
	 * @param TableNode|null $body
	 *
	 * @return void
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userHasCreatedTheFollowingLinkShare(string $user, TableNode  $body): void {
		$response = $this->createLinkShare($user, $body);
		$this->featureContext->theHTTPStatusCodeShouldBe(200, "Failed while creating public share link!", $response);
		$this->featureContext->shareNgAddToCreatedLinkShares($response);
	}

	/**
	 * @When /^user "([^"]*)" updates the last public link share using the Graph API with$/
	 *
	 * @param string $user
	 * @param TableNode|null $body
	 *
	 * @return void
	 * @throws Exception
	 */
	public function userUpdatesLastPublicLinkShareUsingTheGraphApiWith(string $user, TableNode  $body):void {
		$bodyRows = $body->getRowsHash();
		$space = $bodyRows['space'];
		$resourceType = $bodyRows['resourceType'];
		$resource = $bodyRows['resource'];
		$spaceId = ($this->spacesContext->getSpaceByName($user, $space))["id"];
		if ($resourceType === 'folder') {
			$itemId = $this->spacesContext->getResourceId($user, $space, $resource);
		} else {
			$itemId = $this->spacesContext->getFileId($user, $space, $resource);
		}

		if (\array_key_exists('role', $bodyRows) && \array_key_exists('expirationDateTime', $bodyRows)) {
			$body = [
				"expirationDateTime" => $bodyRows['expirationDateTime'],
				"link" => [
					"type" => $bodyRows['permissionsRole']
				]
			];
		} elseif (\array_key_exists('permissionsRole', $bodyRows)) {
			$body = [
				"link" => [
					"type" => $bodyRows['permissionsRole']
				]
			];
		} elseif (\array_key_exists('expirationDateTime', $bodyRows)) {
			$body = [
				"expirationDateTime" => $bodyRows['expirationDateTime']
			];
		} else {
			throw new Error('Expiration date or role is missing to update for share link!');
		}

		$response = GraphHelper::updateLinkShare(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getStepLineRef(),
			$user,
			$this->featureContext->getPasswordForUser($user),
			$spaceId,
			$itemId,
			\json_encode($body),
			$this->featureContext->shareNgGetLastCreatedLinkShareID()
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @When user :user sets the following password for the last link share using the Graph API:
	 *
	 * @param string $user
	 * @param TableNode|null $body
	 *
	 * @return void
	 * @throws Exception
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userSetsOrUpdatesFollowingPasswordForLastLinkShareUsingTheGraphApi(string $user, TableNode  $body):void {
		$bodyRows = $body->getRowsHash();
		$space = $bodyRows['space'];
		$resourceType = $bodyRows['resourceType'];
		$resource = $bodyRows['resource'];
		$spaceId = ($this->spacesContext->getSpaceByName($user, $space))["id"];
		if ($resourceType === 'folder') {
			$itemId = $this->spacesContext->getResourceId($user, $space, $resource);
		} else {
			$itemId = $this->spacesContext->getFileId($user, $space, $resource);
		}

		if (\array_key_exists('password', $bodyRows)) {
			$body = [
				"password" => $this->featureContext->getActualPassword($bodyRows['password']),
			];
		} else {
			throw new Error('Password is missing to set for share link!');
		}

		$response = GraphHelper::setLinkSharePassword(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getStepLineRef(),
			$user,
			$this->featureContext->getPasswordForUser($user),
			$spaceId,
			$itemId,
			\json_encode($body),
			$this->featureContext->shareNgGetLastCreatedLinkShareID()
		);
		$this->featureContext->setResponse($response);
	}

	/**
	 * @param string $sharer
	 * @param string $shareType (user|group)
	 * @param string $resourceType
	 * @param string $resource
	 * @param string $space
	 * @param string|null $sharee can be both user or group
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 * @throws JsonException
	 */
	public function removeSharePermission(
		string $sharer,
		string $shareType,
		string $resourceType,
		string $resource,
		string $space,
		?string $sharee = null
	): ResponseInterface {
		$spaceId = ($this->spacesContext->getSpaceByName($sharer, $space))["id"];
		$itemId = ($resourceType === 'folder')
			? $this->spacesContext->getResourceId($sharer, $space, $resource)
			: $this->spacesContext->getFileId($sharer, $space, $resource);

		$permId = ($shareType === 'link')
			? $this->featureContext->shareNgGetLastCreatedLinkShareID()
			: $this->featureContext->shareNgGetLastCreatedUserGroupShareID();
		return
			GraphHelper::deleteSharePermission(
				$this->featureContext->getBaseUrl(),
				$this->featureContext->getStepLineRef(),
				$sharer,
				$this->featureContext->getPasswordForUser($sharer),
				$spaceId,
				$itemId,
				$permId
			);
	}

	/**
	 * @When /^user "([^"]*)" removes the share permission of (user|group) "([^"]*)" from (file|folder) "([^"]*)" of space "([^"]*)" using the Graph API$/
	 *
	 * @param string $sharer
	 * @param string $shareType (user|group)
	 * @param string $sharee can be both user or group
	 * @param string $resourceType
	 * @param string $resource
	 * @param string $space
	 *
	 * @return void
	 * @throws JsonException
	 * @throws \GuzzleHttp\Exception\GuzzleException
	 */
	public function userRemovesSharePermissionOfUserFromResourceOfSpaceUsingGraphAPI(
		string $sharer,
		string $shareType,
		string $sharee,
		string $resourceType,
		string $resource,
		string $space
	): void {
		$this->featureContext->setResponse(
			$this->removeSharePermission($sharer, $shareType, $resourceType, $resource, $space)
		);
	}

	/**
	 * @When /^user "([^"]*)" removes the share permission of link from (file|folder) "([^"]*)" of space "([^"]*)" using the Graph API$/
	 *
	 * @param string $sharer
	 * @param string $resourceType
	 * @param string $resource
	 * @param string $space
	 *
	 * @return void
	 * @throws JsonException
	 * @throws GuzzleException
	 */
	public function userRemovesSharePermissionOfAResourceInLinkShareUsingGraphAPI(
		string $sharer,
		string $resourceType,
		string $resource,
		string $space
	):void {
		$this->featureContext->setResponse(
			$this->removeSharePermission($sharer, 'link', $resourceType, $resource, $space)
		);
	}

	/**
	 * @Then /^for user "([^"]*)" the space Shares should (not|)\s?contain these (files|entries):$/
	 *
	 * @param string $user
	 * @param string $shouldOrNot
	 * @param TableNode $table
	 *
	 * @return void
	 * @throws Exception
	 */
	public function forUserTheSpaceSharesShouldContainTheseEntries(string $user, string $shouldOrNot, TableNode $table): void {
		$should = $shouldOrNot !== 'not';
		$rows = $table->getRows();
		$response = GraphHelper::getSharesSharedWithMe(
			$this->featureContext->getBaseUrl(),
			$this->featureContext->getStepLineRef(),
			$user,
			$this->featureContext->getPasswordForUser($user)
		);
		$contents = \json_decode($response->getBody()->getContents(), true);

		$fileFound = empty(array_diff(array_map(fn ($row) => trim($row[0], '/'), $rows), array_column($contents['value'], 'name')));

		$assertMessage = $should
			? "Response does not contain the entry."
			: "Response does contain the entry but should not.";

		Assert::assertSame($should, $fileFound, $assertMessage);
	}
}
