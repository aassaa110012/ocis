@api @files_sharing-app-required
Feature: capabilities

  Background:
    Given using OCS API version "1"

  @smokeTest @issue-ocis-1285
  Scenario: getting default capabilities with admin user
    When the administrator retrieves the capabilities using the capabilities API
    Then the OCS status code should be "100"
    And the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "files_sharing"
          ],
          "properties": {
            "files_sharing": {
              "type": "object",
              "required": [
                "user"
              ],
              "properties": {
                "user": {
                  "type": "object",
                  "required": [
                    "profile_picture"
                  ],
                  "properties": {
                    "profile_picture": {
                      "type": "boolean",
                      "enum": [
                        true
                      ]
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
    """

  @files_trashbin-app-required @skipOnReva
  Scenario: getting trashbin app capability with admin user
    When the administrator retrieves the capabilities using the capabilities API
    Then the OCS status code should be "100"
    And the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "files"
          ],
          "properties": {
            "files": {
              "type": "object",
              "required": [
                "undelete"
              ],
              "properties": {
                "undelete": {
                  "type": "boolean",
                  "enum": [
                    true
                  ]
                }
              }
            }
          }
        }
      }
    }
    """

  @files_versions-app-required @skipOnReva
  Scenario: getting versions app capability with admin user
    When the administrator retrieves the capabilities using the capabilities API
    Then the OCS status code should be "100"
    And the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "files"
          ],
          "properties": {
            "files": {
              "type": "object",
              "required": [
                "versioning"
              ],
              "properties": {
                "versioning": {
                  "type": "boolean",
                  "enum": [
                    true
                  ]
                }
              }
            }
          }
        }
      }
    }
    """

  @issue-ocis-1285
  Scenario: getting default_permissions capability with admin user
    When the administrator retrieves the capabilities using the capabilities API
    Then the OCS status code should be "100"
    And the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "files_sharing"
          ],
          "properties": {
            "files_sharing": {
              "type": "object",
              "required": [
                "default_permissions"
              ],
              "properties": {
                "default_permissions": {
                  "type": "number",
                  "enum": [
                    31
                  ]
                }
              }
            }
          }
        }
      }
    }
    """

  @issue-ocis-1285
  Scenario: .htaccess is reported as a blacklisted file by default
    When the administrator retrieves the capabilities using the capabilities API
    Then the OCS status code should be "100"
    And the HTTP status code should be "200"
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "files"
          ],
          "properties": {
            "files": {
              "type": "object",
              "required": [
                "blacklisted_files"
              ],
              "properties": {
                "blacklisted_files": {
                  "type": "array",
                  "minItems": 1,
                  "items": {
                    "type": "string",
                    "enum": [".htaccess"]
                  }
                }
              }
            }
          }
        }
      }
    }
    """

  @smokeTest
  Scenario: getting default capabilities with admin user
    When the administrator retrieves the capabilities using the capabilities API
    And the JSON data of the response should match
    """
    {
      "type": "object",
      "required": [
        "capabilities",
        "version"
      ],
      "properties": {
        "capabilities": {
          "type": "object",
          "required": [
            "core"
          ],
          "properties": {
            "core": {
              "type": "object",
              "required": [
                "status"
              ],
              "properties": {
                "status": {
                  "type": "object",
                  "required": [
                    "edition",
                    "product",
                    "productname",
                    "version",
                    "versionstring"
                  ],
                  "properties": {
                    "edition": {
                      "type": "string",
                      "enum": ["%edition%"]
                    },
                    "product": {
                      "type": "string",
                      "enum": ["%productname%"]
                    },
                    "productname": {
                      "type": "string",
                      "enum": ["%productname%"]
                    },
                    "version": {
                      "type": "string",
                      "enum": ["%version%"]
                    },
                    "versionstring": {
                      "type": "string",
                      "enum": ["%versionstring%"]
                    }
                  }
                }
              }
            }
          }
        },
        "version": {
          "type": "object",
          "required": [
            "string",
            "edition",
            "product"
          ],
          "properties": {
            "string": {
              "type": "string",
              "enum": ["%versionstring%"]
            },
            "edition": {
              "type": "string",
              "enum": ["%edition%"]
            },
            "product": {
              "type": "string",
              "enum": ["%productname%"]
            }
          }
        }
      }
    }
    """
    And the major-minor-micro version data in the response should match the version string
