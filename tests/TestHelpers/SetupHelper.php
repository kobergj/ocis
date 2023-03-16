<?php declare(strict_types=1);
/**
 * ownCloud
 *
 * @author Artur Neumann <artur@jankaritech.com>
 * @copyright Copyright (c) 2017 Artur Neumann artur@jankaritech.com
 *
 * This code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License,
 * as published by the Free Software Foundation;
 * either version 3 of the License, or any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>
 *
 */
namespace TestHelpers;

use Behat\Testwork\Hook\Scope\HookScope;
use GuzzleHttp\Exception\GuzzleException;
use GuzzleHttp\Exception\ServerException;
use Exception;
use Psr\Http\Message\ResponseInterface;
use SimpleXMLElement;

/**
 * Helper to setup UI / Integration tests
 *
 * @author Artur Neumann <artur@jankaritech.com>
 *
 */
class SetupHelper extends \PHPUnit\Framework\Assert {
	/**
	 * @var string
	 */
	private static $ocPath = null;
	/**
	 * @var string
	 */
	private static $baseUrl = null;
	/**
	 * @var string
	 */
	private static $adminUsername = null;
	/**
	 * @var string
	 */
	private static $adminPassword = null;

	/**
	 *
	 * @param HookScope $scope
	 *
	 * @return array of suite context parameters
	 */
	public static function getSuiteParameters(HookScope $scope):array {
		return $scope->getEnvironment()->getSuite()
			->getSettings() ['context'] ['parameters'];
	}

	/**
	 * Fixup OC path so that it always starts with a "/" and does not end with
	 * a "/".
	 *
	 * @param string|null $ocPath
	 *
	 * @return string
	 */
	private static function normaliseOcPath(?string $ocPath):string {
		return '/' . \trim($ocPath, '/');
	}

	/**
	 *
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return void
	 */
	public static function init(
		?string $adminUsername,
		?string $adminPassword,
		?string $baseUrl,
		?string $ocPath
	): void {
		foreach (\func_get_args() as $variableToCheck) {
			if (!\is_string($variableToCheck)) {
				throw new \InvalidArgumentException(
					"mandatory argument missing or wrong type ($variableToCheck => "
					. \gettype($variableToCheck) . ")"
				);
			}
		}
		self::$adminUsername = $adminUsername;
		self::$adminPassword = $adminPassword;
		self::$baseUrl = \rtrim($baseUrl, '/');
		self::$ocPath = self::normaliseOcPath($ocPath);
	}

	/**
	 *
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $xRequestId
	 *
	 * @return SimpleXMLElement
	 * @throws GuzzleException
	 */
	public static function getSysInfo(
		?string $baseUrl,
		?string $adminUsername,
		?string $adminPassword,
		?string $xRequestId = ''
	):SimpleXMLElement {
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"GET",
			"/apps/testing/api/v1/sysinfo",
			$xRequestId
		);
		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not get sysinfo " . $result->getReasonPhrase()
			);
		}
		return HttpRequestHelper::getResponseXml($result, __METHOD__)->data;
	}

	/**
	 *
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $xRequestId
	 *
	 * @return string
	 * @throws GuzzleException
	 */
	public static function getServerRoot(
		?string $baseUrl,
		?string $adminUsername,
		?string $adminPassword,
		?string $xRequestId = ''
	):string {
		$sysInfo = self::getSysInfo(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			$xRequestId
		);
		// server_root is a SimpleXMLElement object that "wraps" a string.
		/// We want the bare string.
		return (string) $sysInfo->server_root;
	}

	/**
	 * @param string|null $adminUsername
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkAdminUsername(?string $adminUsername, ?string $callerName):?string {
		if (self::$adminUsername === null
			&& $adminUsername === null
		) {
			throw new Exception(
				"$callerName called without adminUsername - pass the username or call SetupHelper::init"
			);
		}
		if ($adminUsername === null) {
			$adminUsername = self::$adminUsername;
		}
		return $adminUsername;
	}

	/**
	 * @param string|null $adminPassword
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkAdminPassword(?string $adminPassword, ?string $callerName):string {
		if (self::$adminPassword === null
			&& $adminPassword === null
		) {
			throw new Exception(
				"$callerName called without adminPassword - pass the password or call SetupHelper::init"
			);
		}
		if ($adminPassword === null) {
			$adminPassword = self::$adminPassword;
		}
		return $adminPassword;
	}

	/**
	 * @param string|null $baseUrl
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkBaseUrl(?string $baseUrl, ?string $callerName):?string {
		if (self::$baseUrl === null
			&& $baseUrl === null
		) {
			throw new Exception(
				"$callerName called without baseUrl - pass the baseUrl or call SetupHelper::init"
			);
		}
		if ($baseUrl === null) {
			$baseUrl = self::$baseUrl;
		}
		return $baseUrl;
	}

	/**
	 *
	 * @param string|null $dirPathFromServerRoot e.g. 'apps2/myapp/appinfo'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function mkDirOnServer(
		?string $dirPathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "mkDirOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "mkDirOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "mkDirOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"POST",
			"/apps/testing/api/v1/dir",
			$xRequestId,
			['dir' => $dirPathFromServerRoot]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not create directory $dirPathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $dirPathFromServerRoot e.g. 'apps2/myapp/appinfo'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function rmDirOnServer(
		?string $dirPathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "rmDirOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "rmDirOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "rmDirOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"DELETE",
			"/apps/testing/api/v1/dir",
			$xRequestId,
			['dir' => $dirPathFromServerRoot]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not delete directory $dirPathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $filePathFromServerRoot e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $content
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 */
	public static function createFileOnServer(
		?string $filePathFromServerRoot,
		?string $content,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "createFileOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "createFileOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "createFileOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"POST",
			"/apps/testing/api/v1/file",
			$xRequestId,
			[
				'file' => $filePathFromServerRoot,
				'content' => $content
			]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not create file $filePathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $filePathFromServerRoot e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function deleteFileOnServer(
		?string $filePathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "deleteFileOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "deleteFileOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "deleteFileOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"DELETE",
			"/apps/testing/api/v1/file",
			$xRequestId,
			[
				'file' => $filePathFromServerRoot
			]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not delete file $filePathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 * @param string|null $fileInCore e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return string
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function readFileFromServer(
		?string $fileInCore,
		?string $xRequestId = '',
		?string $baseUrl  = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):string {
		$baseUrl = self::checkBaseUrl($baseUrl, "readFile");
		$adminUsername = self::checkAdminUsername(
			$adminUsername,
			"readFile"
		);
		$adminPassword = self::checkAdminPassword(
			$adminPassword,
			"readFile"
		);

		$response = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			'GET',
			"/apps/testing/api/v1/file?file={$fileInCore}",
			$xRequestId
		);
		self::assertSame(
			200,
			$response->getStatusCode(),
			"Failed to read the file {$fileInCore}"
		);
		$localContent = HttpRequestHelper::getResponseXml($response, __METHOD__);
		$localContent = (string)$localContent->data->element->contentUrlEncoded;
		return \urldecode($localContent);
	}

	/**
	 * returns the content of a file in a skeleton folder
	 *
	 * @param string|null $fileInSkeletonFolder
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return string content of the file
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function readSkeletonFile(
		?string $fileInSkeletonFolder,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):string {
		$baseUrl = self::checkBaseUrl($baseUrl, "readSkeletonFile");
		$adminUsername = self::checkAdminUsername(
			$adminUsername,
			"readSkeletonFile"
		);
		$adminPassword = self::checkAdminPassword(
			$adminPassword,
			"readSkeletonFile"
		);

		$fileInSkeletonFolder = \rawurlencode("/$fileInSkeletonFolder");
		$response = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			'GET',
			"/apps/testing/api/v1/file?file={$fileInSkeletonFolder}&absolute=true",
			$xRequestId
		);
		self::assertSame(
			200,
			$response->getStatusCode(),
			"Failed to read the file {$fileInSkeletonFolder}"
		);
		$localContent = HttpRequestHelper::getResponseXml($response, __METHOD__);
		$localContent = (string)$localContent->data->element->contentUrlEncoded;
		$localContent = \urldecode($localContent);
		return $localContent;
	}
}
