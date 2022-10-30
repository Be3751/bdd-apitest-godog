Feature: get address
    In order to know address from postal code
    As an API user
    I need to be able to request address

    Scenario Outline: 
        When I send request with the query param <zipcode>
        Then the response code should be 200 
        And the fields of the response JSON should meet the restriction:
        And the response JSON should match the schema:
        And the response JSON should match the JSON file <filename>

        Examples:
            | zipcode | filename |
            | "6810002" | "address01.json" |
            | "5110206" | "address02.json" |
            | "8952201" | "address03.json" |