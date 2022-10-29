# file: version.feature
Feature: get version
  In order to know godog version
  As an API user
  I need to be able to request version

  Scenario Outline: does not allow methods other than GET method
    When I send <Method> request to "/version"
    Then the response code should be 405
    And the response should match json:
      """
      {
        "error": "Method not allowed"
      }
      """

    Examples:
      | Method |
      | "POST" |
      | "PUT" |
      | "DELETE" |

  Scenario: should get version number
    When I send "GET" request to "/version"
    Then the response code should be 200
    And the response should match json:
      """
      {
        "version": "v0.0.0-dev"
      }
      """