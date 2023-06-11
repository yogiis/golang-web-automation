package stepdefinitions

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/yogiis/golang-web-automation/helpers"
)

func InitializeScenario(sc *godog.ScenarioContext) {

	step := &Entity{}

	sc.Step(`I open website "([^"]*)"`, step.IOpenTheWebsite)
	sc.Step(`I fill form in following information:`, step.IFillForm)
	sc.Step(`I click submit button`, step.IClickSubmit)
	sc.Step(`Verify result information "([^"]*)"`, step.VerifyResult)

	sc.After(func(ctx context.Context, _ *godog.Scenario, err error) (context.Context, error) {
		if err := step.Page.Close(); err != nil {
			helpers.LogPanicln(err)
		}
		if err := step.Browser.Close(); err != nil {
			helpers.LogPanicln(err)
		}
		if err := step.Pw.Stop(); err != nil {
			helpers.LogPanicln(err)
		}
		return ctx, nil
	})
}
