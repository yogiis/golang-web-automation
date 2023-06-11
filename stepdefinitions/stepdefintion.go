package stepdefinitions

import (
	"strings"

	"github.com/cucumber/godog"
	"github.com/playwright-community/playwright-go"
	"github.com/yogiis/golang-web-automation/helpers"
)

func (e *Entity) IOpenTheWebsite(params string) error {
	pw, err := playwright.Run()
	helpers.LogPanicln(err)
	e.Pw = pw

	browser, err := pw.Firefox.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	helpers.LogPanicln(err)
	e.Browser = browser

	page, err := browser.NewPage()
	helpers.LogPanicln(err)

	err = page.SetViewportSize(1920, 1440)
	helpers.LogPanicln(err)

	e.Page = page

	_, err = page.Goto(params)
	helpers.LogPanicln(err)
	return nil
}

func (e *Entity) IFillForm(table *godog.Table) error {

	for _, row := range table.Rows {
		field := row.Cells[0]
		value := row.Cells[1]

		switch field.Value {
		case "Username":
			userName, err := e.Page.Locator("#userName")
			helpers.LogPanicln(err)

			err = userName.Fill(value.Value)
			helpers.LogPanicln(err)

			fullname, err := userName.TextContent()
			helpers.LogPanicln(err)

			e.FullName = fullname
		case "Email":
			userEmail, err := e.Page.Locator("#userEmail")
			helpers.LogPanicln(err)

			err = userEmail.Fill(value.Value)
			helpers.LogPanicln(err)
		case "Current Address":
			currentAddress, err := e.Page.Locator("#currentAddress")
			helpers.LogPanicln(err)

			err = currentAddress.Fill(value.Value)
			helpers.LogPanicln(err)
		case "Permanent Address":
			permanentAddress, err := e.Page.Locator("#permanentAddress")
			helpers.LogPanicln(err)

			err = permanentAddress.Fill(value.Value)
			helpers.LogPanicln(err)
		}

	}
	return nil
}

func (e *Entity) IClickSubmit() error {
	submitButton, err := e.Page.Locator("#submit")
	helpers.LogPanicln(err)
	err = submitButton.Click()
	helpers.LogPanicln(err)

	return nil
}

func (e *Entity) VerifyResult(expected string) error {
	nameView, err := e.Page.Locator("#name")
	helpers.LogPanicln(err)

	name, err := nameView.TextContent()
	actual := strings.TrimPrefix(name, "Name:")
	helpers.LogPanicln(err)

	e.Cases.AssertEqual(expected, actual)

	return nil
}
