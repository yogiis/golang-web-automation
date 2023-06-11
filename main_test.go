package main

import (
	"flag"
	"testing"

	"github.com/cucumber/godog"
	suiteTest "github.com/yogiis/golang-web-automation/stepdefinitions"
)

var opts = godog.Options{}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {

	opts := godog.Options{
		Format:   "pretty,cucumber:./report/cucumber.json",
		Paths:    []string{"features"},
		TestingT: t,
	}

	flag.Parse()

	suite := godog.TestSuite{
		ScenarioInitializer: suiteTest.InitializeScenario,
		Options:             &opts,
	}

	if suite.Run() != 0 {
		t.Fatal("Error")
	}

}
