package main

import (
	"fmt"

	"gitlab.com/rodrigoodhin/gocure/models"
	"gitlab.com/rodrigoodhin/gocure/pkg/gocure"
	"gitlab.com/rodrigoodhin/gocure/report/html"
)

func main() {

	// Set html options
	html := gocure.HTML{
		Config: html.Data{
			InputJsonPath:     "./report/cucumber.json",
			OutputHtmlFolder:  "./report/",
			ShowEmbeddedFiles: true,
			Metadata: models.Metadata{
				AppVersion:      "0.1.0",
				TestEnvironment: "development",
				Browser:         "Google Chrome",
				Platform:        "Linux",
				Parallel:        "Scenarios",
				Executed:        "Remote",
			},
		},
	}

	// Generate HTML report
	err := html.Generate()
	if err != nil {
		fmt.Printf("error generatig html report: %v", err)
	}

}
