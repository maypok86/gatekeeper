Feature: Gatekeeper
  In order to use Gatekeeper
  As an GRPC client
  I need to be able to send grpc requests

  Scenario Outline: should allow for Gatekeeper
    Given login is "<login>"
    And password is "<password>"
    And ip is "<ip>"
    When I call grpc method "Allow"
    Then response error should be "<error>"

    Examples:
      | login  | password | ip      | error |
      | login1 | pass1    | 1.2.3.4 | nil   |
      | login1 | pass1    | 1.2.3.4 | nil   |
      | login1 | pass1    | 1.2.3.4 | nil   |
      | login1 | pass2    | 1.2.3.4 | nil   |

  Scenario Outline: should reset login from bucket
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method      | login  |  ip      | password | error |
      | Allow       | login1 | 1.2.3.4  | pass2    | nil   |
      | ResetLogin  | login1 |          |          | nil   |
      | Allow       | login1 | 1.2.3.4  | pass2    | nil   |

  Scenario Outline: should reset ip from bucket
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method  | login  | ip      | password | error |
      | Allow   | login1 | 1.2.3.4 | pass2    | nil   |
      | ResetIP |        | 1.2.3.4 |          | nil   |
      | Allow   | login1 | 1.2.3.4 | pass2    | nil   |

  Scenario Outline: should add subnet to blacklist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method       | login  | password | ip           | error |
      | Allow        | login3 | pass3    | 192.168.0.30 | nil   |
      | BlacklistAdd |        |          |              | nil   |
      | Allow        | login3 | pass3    | 192.168.0.30 | nil   |

  Scenario Outline: should remove subnet from blacklist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method          | login  | password | ip           | error |
      | Allow           | login4 | pass4    | 192.168.0.30 | nil   |
      | BlacklistRemove |        |          |              | nil   |
      | Allow           | login4 | pass4    | 192.168.0.30 | nil   |

  Scenario Outline: should add subnet to whitelist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method       | login  | password | ip           | error |
      | Allow        | login5 | pass5    | 192.168.0.30 | nil   |
      | Allow        | login5 | pass5    | 192.168.0.30 | nil   |
      | Allow        | login5 | pass5    | 192.168.0.30 | nil   |
      | WhitelistAdd |        |          |              | nil   |
      | Allow        | login5 | pass5    | 192.168.0.30 | nil   |

  Scenario Outline: should remove subnet from whitelist
    Given login is "<login>"
    And ip is "<ip>"
    And password is "<password>"
    And subnet is "192.168.0.0/25"

    When I call grpc method "<method>"
    Then response error should be "<error>"

    Examples:
      | method          | login  | password | ip           | error |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |
      | WhitelistRemove |        |          |              | nil   |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |
      | Allow           | login6 | pass6    | 192.168.0.30 | nil   |