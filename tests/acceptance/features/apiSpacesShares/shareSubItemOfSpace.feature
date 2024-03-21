Feature: Share a file or folder that is inside a space
  As a user with manager space role
  I want to be able to share the data inside the space
  So that other users can have access to it

  Note - this feature is run in CI with ACCOUNTS_HASH_DIFFICULTY set to the default for production
  See https://github.com/owncloud/ocis/issues/1542 and https://github.com/owncloud/ocis/pull/839

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
      | Bob      |
    And using spaces DAV path
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "share sub-item" with the default quota using the Graph API
    And user "Alice" has created a folder "folder" in space "share sub-item"
    And user "Alice" has uploaded a file inside space "share sub-item" with content "some content" to "file.txt"
    And using new DAV path


  Scenario Outline: manager of the space can share an entity inside project space to another user with role
    When user "Alice" creates a share inside of space "share sub-item" with settings:
      | path       | <resource>    |
      | shareWith  | Brian         |
      | role       | <space-role>  |
      | expireDate | <expire-date> |
    Then the HTTP status code should be "200"
    And the OCS status code should be "200"
    And the OCS status message should be "OK"
    And as "Brian" <resource-type> "Shares/<resource>" should exist
    And the information about the last share for user "Brian" should include
      | expiration | <expiration> |
    Examples:
      | resource | resource-type | space-role | expire-date              | expiration |
      | folder   | folder        | viewer     |                          |            |
      | folder   | folder        | editor     | 2042-03-25T23:59:59+0100 | 2042-03-25 |
      | file.txt | file          | viewer     |                          |            |
      | file.txt | file          | editor     | 2042-03-25T23:59:59+0100 | 2042-03-25 |


  Scenario Outline: user participant of the project space with manager role can share an entity to another user
    Given user "Alice" has shared a space "share sub-item" with settings:
      | shareWith | Brian   |
      | role      | manager |
    When user "Brian" creates a share inside of space "share sub-item" with settings:
      | path       | <resource>    |
      | shareWith  | Bob           |
      | role       | <space-role>  |
      | expireDate | <expire-date> |
    Then the HTTP status code should be "200"
    And the OCS status code should be "200"
    And the OCS status message should be "OK"
    And as "Bob" <resource-type> "Shares/<resource>" should exist
    And the information about the last share for user "Brian" should include
      | expiration | <expiration> |
    Examples:
      | resource | resource-type | space-role | expire-date              | expiration |
      | folder   | folder        | viewer     | 2042-03-25T23:59:59+0100 | 2042-03-25 |
      | folder   | folder        | editor     |                          |            |
      | file.txt | file          | viewer     | 2042-03-25T23:59:59+0100 | 2042-03-25 |
      | file.txt | file          | editor     |                          |            |

  @skipOnRevaMaster
  Scenario Outline: user participant of the project space without space manager role cannot share an entity to another user
    Given user "Alice" has shared a space "share sub-item" with settings:
      | shareWith | Brian        |
      | role      | <space-role> |
    When user "Brian" creates a share inside of space "share sub-item" with settings:
      | path      | <resource> |
      | shareWith | Bob        |
      | role      | editor     |
    Then the HTTP status code should be "403"
    And the OCS status code should be "403"
    And the OCS status message should be "No share permission"
    Examples:
      | resource | space-role |
      | folder   | editor     |
      | file.txt | editor     |
      | file.txt | viewer     |
      | folder   | viewer     |


  Scenario Outline: user participant of the project space can see the created resources share
    Given user "Alice" has shared a space "share sub-item" with settings:
      | shareWith | Brian        |
      | role      | <space-role> |
    When user "Alice" creates a share inside of space "share sub-item" with settings:
      | path      | file.txt |
      | shareWith | Bob      |
      | role      | editor   |
    Then for user "Alice" the space "share sub-item" should contain the last created share of the file "file.txt"
    And for user "Brian" the space "share sub-item" should contain the last created share of the file "file.txt"
    Examples:
      | space-role |
      | editor     |
      | viewer     |
      | manager    |


  Scenario: user shares the folder to the group
    Given group "sales" has been created
    And the administrator has added a user "Brian" to the group "sales" using the Graph API
    When user "Alice" creates a share inside of space "share sub-item" with settings:
      | path       | folder                   |
      | shareWith  | sales                    |
      | shareType  | 1                        |
      | role       | viewer                   |
      | expireDate | 2042-01-01T23:59:59+0100 |
    Then the HTTP status code should be "200"
    And the OCS status code should be "200"
    And the OCS status message should be "OK"
    And as "Brian" folder "Shares/folder" should exist
    And the information about the last share for user "Brian" should include
      | expiration | 2042-01-01 |


  Scenario: user changes the expiration date
    Given user "Alice" has created a share inside of space "share sub-item" with settings:
      | path       | folder                   |
      | shareWith  | Brian                    |
      | role       | viewer                   |
      | expireDate | 2042-01-01T23:59:59+0100 |
    When user "Alice" changes the last share with settings:
      | expireDate | 2044-01-01T23:59:59.999+01:00 |
    Then the HTTP status code should be "200"
    And the information about the last share for user "Brian" should include
      | expiration | 2044-01-01 |


  Scenario: user deletes the expiration date
    Given user "Alice" has created a share inside of space "share sub-item" with settings:
      | path       | folder                   |
      | shareWith  | Brian                    |
      | role       | viewer                   |
      | expireDate | 2042-01-01T23:59:59+0100 |
    When user "Alice" changes the last share with settings:
      | expireDate |  |
    Then the HTTP status code should be "200"
    And the information about the last share for user "Brian" should include
      | expiration |  |


  Scenario: check the end of expiration date in user share
    Given user "Alice" has created a share inside of space "share sub-item" with settings:
      | path       | folder                   |
      | shareWith  | Brian                    |
      | role       | viewer                   |
      | expireDate | 2042-01-01T23:59:59+0100 |
    When user "Alice" expires the last share
    Then the HTTP status code should be "200"
    And as "Brian" folder "Shares/folder" should not exist

  @issue-5823
  Scenario: check the end of expiration date in group share
    Given group "sales" has been created
    And the administrator has added a user "Brian" to the group "sales" using the Graph API
    And user "Alice" has created a share inside of space "share sub-item" with settings:
      | path       | folder                   |
      | shareWith  | sales                    |
      | shareType  | 1                        |
      | role       | viewer                   |
      | expireDate | 2042-01-01T23:59:59+0100 |
    When user "Alice" expires the last share
    Then the HTTP status code should be "200"
    And as "Brian" folder "Shares/folder" should not exist

  @issue-enterprise-6423 @env-config
  Scenario Outline: user cannot share items in the project space with share permission if resharing is denied
    Given the config "OCIS_ENABLE_RESHARING" has been set to "false"
    And user "Alice" has shared a space "share sub-item" with settings:
      | shareWith | Brian  |
      | role      | viewer |
    When user "Alice" creates a share inside of space "share sub-item" with settings:
      | path        | folder        |
      | shareWith   | Bob           |
      | role        | custom        |
      | permissions | <permissions> |
    Then the HTTP status code should be "400"
    And the OCS status code should be "400"
    And the OCS status message should be "resharing not supported"
    Examples:
      | permissions | description                  |
      | 19          | view + edit                  |
      | 21          | view + create                |
      | 23          | view + create + edit         |
      | 25          | view + delete                |
      | 27          | view + edit + delete         |
      | 29          | view + create + delete       |
      | 31          | view + create + edit +delete |


  @issue-enterprise-6423 @env-config
  Scenario Outline: user cannot share items in the personal space with share permission if resharing is denied
    Given the config "OCIS_ENABLE_RESHARING" has been set to "false"
    And user "Alice" has uploaded file with content "some content" to "/file.txt"
    When user "Alice" creates a share inside of space "Alice Hansen" with settings:
      | path        | file.txt      |
      | shareWith   | Bob           |
      | role        | custom        |
      | permissions | <permissions> |
    Then the HTTP status code should be "400"
    And the OCS status code should be "400"
    And the OCS status message should be "resharing not supported"
    Examples:
      | permissions | description                  |
      | 19          | view + edit                  |
      | 21          | view + create                |
      | 23          | view + create + edit         |
      | 25          | view + delete                |
      | 27          | view + edit + delete         |
      | 29          | view + create + delete       |
      | 31          | view + create + edit +delete |
