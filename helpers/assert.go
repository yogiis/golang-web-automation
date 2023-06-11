package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	testing testing.T
}

func (t *Case) assertNew() *assert.Assertions {
	return assert.New(&t.testing)
}

func (t *Case) AssertEqual(expected, actual interface{}) bool {
	equal := t.assertNew().Equal(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual))
	if !equal {
		LogPanicln(ErrorHandleEqual(expected, actual))
	}
	return true
}
