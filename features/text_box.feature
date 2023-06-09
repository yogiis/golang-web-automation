Feature: Text Box

    Scenario: Ensure user successfully submitted form personal data
        Given I open website "https://demoqa.com/text-box"
        When I fill form in following information:
            | Field             | Value                |
            | Username          | Yogi                 |
            | Email             | yogiis.ari@gmail.com |
            | Current Address   | Yogyakarta           |
            | Permanent Address | Semarang             |
        And I click submit button
        Then Verify result information "Yogi"

    Scenario: Ensure user successfully submitted form personal data
        Given I open website "https://demoqa.com/text-box"
        When I fill form in following information:
            | Field             | Value               |
            | Username          | Ari                 |
            | Email             | Ari047.is@gmail.com |
            | Current Address   | Surabaya            |
            | Permanent Address | Magelang            |
        And I click submit button
        Then Verify result information "Ari"
