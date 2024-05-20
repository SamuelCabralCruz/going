package matcher

import (
	"errors"
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/onsi/gomega/types"
)

type beFunctionCustomMatcher struct {
	ref string
}

func (m *beFunctionCustomMatcher) Match(actual any) (bool, error) {
	if !phi.IsFunction(actual) {
		return false, errors.New("actual is not a function")
	}
	return m.ref == phi.FunctionFullPath(actual), nil
}

func (m *beFunctionCustomMatcher) FailureMessage(actual any) string {
	return fmt.Sprintf("Expected %s to be identical to %s", phi.FunctionFullPath(actual), m.ref)
}

func (m *beFunctionCustomMatcher) NegatedFailureMessage(actual any) string {
	return fmt.Sprintf("Expected %s to be different of %s", phi.FunctionFullPath(actual), m.ref)
}

func BeFunction(f any) types.GomegaMatcher {
	if !phi.IsFunction(f) {
		panic(errors.New("input parameter must be a function"))
	}
	return &beFunctionCustomMatcher{ref: phi.FunctionFullPath(f)}
}
