@ocm
Feature: delete federated connections
  As a user
  I want to delete federated connections if they are no longer needed

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And using server "REMOTE"
    And user "Brian" has been created with default attributes and without skeleton files


  Scenario: federated user deletes the federated connection
    Given using server "LOCAL"
    And "Alice" has created the federation share invitation
    And using server "REMOTE"
    And "Brian" has accepted invitation
    When user "Brian" deletes federated connection with user "Alice" using the Graph API
    Then the HTTP status code should be "200"
    When user "Brian" searches for federated user "alice" using Graph API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "minItems": 0,
            "maxItems": 0
          }
        }
      }
      """

  @issue-10216
  Scenario: local user should not be able to find federated user after federated user has deleted connection
    Given using server "LOCAL"
    And "Alice" has created the federation share invitation
    And using server "REMOTE"
    And "Brian" has accepted invitation
    And user "Brian" has deleted federated connection with user "Alice"
    And using server "LOCAL"
    When user "Alice" searches for federated user "brian" using Graph API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "minItems": 0,
            "maxItems": 0
          }
        }
      }
      """

  @issue-10216
  Scenario: federated user should not be able to find federated share after federated user has deleted connection
    Given using server "LOCAL"
    And "Alice" has created the federation share invitation
    And using server "REMOTE"
    And "Brian" has accepted invitation
    And using server "LOCAL"
    And user "Alice" has created folder "folderToShare"
    And user "Alice" has sent the following resource share invitation to federated user:
      | resource        | folderToShare                 |
      | space           | Personal                      |
      | sharee          | Brian                         |
      | shareType       | user                          |
      | permissionsRole | Viewer                        |
      | federatedServer | @federation-ocis-server:10200 |
    And using server "REMOTE"
    When user "Brian" deletes federated connection with user "Alice" using the Graph API
    Then the HTTP status code should be "200"
    When user "Brian" lists the shares shared with him without retry using the Graph API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "minItems": 0,
            "maxItems": 0,
          }
        }
      }
      """

  @issue-10213
  Scenario: federated user should not be able to find federated share after local user has deleted connection
    Given using server "LOCAL"
    And "Alice" has created the federation share invitation
    And using server "REMOTE"
    And "Brian" has accepted invitation
    And using server "LOCAL"
    And user "Alice" has created folder "folderToShare"
    And user "Alice" has sent the following resource share invitation to federated user:
      | resource        | folderToShare                 |
      | space           | Personal                      |
      | sharee          | Brian                         |
      | shareType       | user                          |
      | permissionsRole | Viewer                        |
      | federatedServer | @federation-ocis-server:10200 |
    When user "Alice" deletes federated connection with user "Brian" using the Graph API
    Then the HTTP status code should be "200"
    And using server "REMOTE"
    When user "Brian" lists the shares shared with him without retry using the Graph API
    Then the HTTP status code should be "200"
    And the JSON data of the response should match
      """
      {
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "value": {
            "type": "array",
            "minItems": 0,
            "maxItems": 0,
          }
        }
      }
      """
