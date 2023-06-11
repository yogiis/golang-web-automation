package helpers

import (
	"fmt"
	"log"
)

func LogPanicln(err interface{}) error {
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

func ErrorHandleEqual(expected, actual interface{}) string {
	return fmt.Sprintf("Expected : %v Actual : %v", expected, actual)
}
