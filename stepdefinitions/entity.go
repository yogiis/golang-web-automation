package stepdefinitions

import (
	"github.com/playwright-community/playwright-go"
	"github.com/yogiis/golang-web-automation/helpers"
)

type Entity struct {
	Pw       *playwright.Playwright
	Browser  playwright.Browser
	Page     playwright.Page
	FullName string
	Cases    helpers.Case
}
