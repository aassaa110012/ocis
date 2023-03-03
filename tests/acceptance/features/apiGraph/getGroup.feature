@api @skipOnOcV10
Feature: get groups and their members
  As an admin
  I want to be able to get groups
  So that I can see all the groups and their members

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And the administrator has given "Alice" the role "Admin" using the settings api


  Scenario: admin user lists all the groups
    Given group "tea-lover" has been created
    And group "coffee-lover" has been created
    And group "h2o-lover" has been created
    When user "Alice" gets all the groups using the Graph API
    Then the HTTP status code should be "200"
    And the extra groups returned by the API should be
      | tea-lover    |
      | coffee-lover |
      | h2o-lover    |


  Scenario: normal user cannot get the groups list
    Given user "Brian" has been created with default attributes and without skeleton files
    And group "tea-lover" has been created
    And group "coffee-lover" has been created
    And group "h2o-lover" has been created
    When user "Brian" gets all the groups using the Graph API
    Then the HTTP status code should be "401"
    And the last response should be an unauthorized response


  Scenario: admin user gets users of a group
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Brian    |
      | Carol    |
    And group "tea-lover" has been created
    And user "Brian" has been added to group "tea-lover"
    And user "Carol" has been added to group "tea-lover"
    When user "Alice" gets all the members of group "tea-lover" using the Graph API
    Then the HTTP status code should be "200"
    And the users returned by the API should be
      | Brian |
      | Carol |


  Scenario: normal user tries to get users of a group
    Given user "Brian" has been created with default attributes and without skeleton files
    And group "tea-lover" has been created
    When user "Brian" gets all the members of group "tea-lover" using the Graph API
    Then the HTTP status code should be "401"
    And the last response should be an unauthorized response


  Scenario: admin user gets all groups along with its member's information
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Brian    |
      | Carol    |
    And group "tea-lover" has been created
    And group "coffee-lover" has been created
    And user "Alice" has been added to group "tea-lover"
    And user "Brian" has been added to group "coffee-lover"
    And user "Carol" has been added to group "tea-lover"
    When user "Alice" retrieves all groups along with their members using the Graph API
    Then the HTTP status code should be "200"
    And the JSON data of the response should contain group name "coffee-lover" and match
    """
    {
      "type": "object",
      "required": [
        "members"
      ],
      "properties": {
        "members": {
          "type": "array",
          "items": [
            {
              "type": "object",
              "required": [
                "displayName",
                "id",
                "mail",
                "onPremisesSamAccountName"
              ],
              "properties": {
                "displayName": {
                  "type": "string",
                  "enum": ["Brian Murphy"]
                },
                "id" : {
                  "type": "string",
                  "pattern": "^%uuid_v4%$"
                },
                "mail": {
                  "type": "string",
                  "enum": ["brian@example.org"]
                },
                "onPremisesSamAccountName": {
                  "type": "string",
                  "enum": ["Brian"]
                }
              }
            }
          ]
        }
      }
    }
    """
    And the JSON data of the response should contain group name "tea-lover" and match
    """
    {
      "type": "object",
      "required": [
        "members"
      ],
      "properties": {
        "members": {
          "type": "array",
          "items": [
            {
              "type": "object",
              "required": [
                "displayName",
                "id",
                "mail",
                "onPremisesSamAccountName"
              ],
              "properties": {
                "displayName": {
                  "type": "string",
                  "enum": ["Alice Hansen"]
                },
                "id" : {
                  "type": "string",
                  "pattern": "^%uuid_v4%$"
                },
                "mail": {
                  "type": "string",
                  "enum": ["alice@example.org"]
                },
                "onPremisesSamAccountName": {
                  "type": "string",
                  "enum": ["Alice"]
                }
              }
            },
            {
              "type": "object",
              "required": [
                "displayName",
                "id",
                "mail",
                "onPremisesSamAccountName"
              ],
              "properties": {
                "displayName": {
                  "type": "string",
                  "enum": ["Carol King"]
                },
                "id" : {
                  "type": "string",
                  "pattern": "^%uuid_v4%$"
                },
                "mail": {
                  "type": "string",
                  "enum": ["carol@example.org"]
                },
                "onPremisesSamAccountName": {
                  "type": "string",
                  "enum": ["Carol"]
                }
              }
            }
          ]
        }
      }
    }
    """


  Scenario: normal user gets all groups along with their members information
    Given user "Brian" has been created with default attributes and without skeleton files
    And group "tea-lover" has been created
    And group "coffee-lover" has been created
    And user "Alice" has been added to group "tea-lover"
    And user "Brian" has been added to group "coffee-lover"
    When user "Brian" retrieves all groups along with their members using the Graph API
    Then the HTTP status code should be "401"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "object",
          "required": [
            "message"
          ],
          "properties": {
            "type": "string",
            "enum": ["Unauthorized"]
          }
        }
      }
    }
    """


  Scenario: admin user gets a group along with their members information
    Given user "Brian" has been created with default attributes and without skeleton files
    And group "tea-lover" has been created
    And user "Alice" has been added to group "tea-lover"
    And user "Brian" has been added to group "tea-lover"
    When user "Alice" gets all the members information of group "tea-lover" using the Graph API
    And the JSON data of the response should contain group name "tea-lover" and match
    """
    {
      "type": "object",
      "required": [
        "members"
      ],
      "properties": {
        "members": {
          "type": "array",
          "items": [
            {
              "type": "object",
              "required": [
                "displayName",
                "id",
                "mail",
                "onPremisesSamAccountName"
              ],
              "properties": {
                "displayName": {
                  "type": "string",
                  "enum": ["Alice Hansen"]
                },
                "id" : {
                  "type": "string",
                  "pattern": "^%uuid_v4%$"
                },
                "mail": {
                  "type": "string",
                  "enum": ["alice@example.org"]
                },
                "onPremisesSamAccountName": {
                  "type": "string",
                  "enum": ["Alice"]
                }
              }
            },
            {
              "type": "object",
              "required": [
                "displayName",
                "id",
                "mail",
                "onPremisesSamAccountName"
              ],
              "properties": {
                "displayName": {
                  "type": "string",
                  "enum": ["Brian Murphy"]
                },
                "id" : {
                  "type": "string",
                  "pattern": "^%uuid_v4%$"
                },
                "mail": {
                  "type": "string",
                  "enum": ["brian@example.org"]
                },
                "onPremisesSamAccountName": {
                  "type": "string",
                  "enum": ["Brian"]
                }
              }
            }
          ]
        }
      }
    }
    """

  Scenario: normal user gets a group along with their members information
    Given user "Brian" has been created with default attributes and without skeleton files
    And group "tea-lover" has been created
    And user "Alice" has been added to group "tea-lover"
    And user "Brian" has been added to group "tea-lover"
    When user "Brian" gets all the members information of group "tea-lover" using the Graph API
    Then the HTTP status code should be "401"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "error"
      ],
      "properties": {
        "error": {
          "type": "object",
          "required": [
            "message"
          ],
          "properties": {
            "type": "string",
            "enum": ["Unauthorized"]
          }
        }
      }
    }
    """
