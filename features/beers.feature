Feature: beers
  In order to use beers api
  As an API beer
  I need to be able to manage beers

  Scenario: should get empty beers
    When I send "GET" request to "/beers"
    Then the response code should by 200
    And the response should match json:
    """
      []
    """
    Scenario: should get users
      Given there are users:
        | id | name   | brewery | country | price | currency |
        | 1  | Pilsen | Backus  | Peru    | 5.00  | PEN      |
      When I send "GET" request to "/users"
      Then the response code should be 200
      And the response should match json:
      """
        [
          {
            "id": 1
            "name": "Pilsen"
            "brewery": "Backus"
            "country": "Peru"
            "price": 5.00
            "currency": "PEN"
          }
        ]
      """