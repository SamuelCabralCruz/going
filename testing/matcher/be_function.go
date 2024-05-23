package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/onsi/gomega/types"
)

type beFunctionMatcher struct {
	ref string
}

func (m *beFunctionMatcher) Match(actual any) (bool, error) {
	if !phi.IsFunction(actual) {
		return false, fmt.Errorf("actual must be a function, received `%T`", actual)
	}
	return m.ref == phi.FunctionFullPath(actual), nil
}

func (m *beFunctionMatcher) FailureMessage(actual any) string {
	return fmt.Sprintf("expected %s to be identical to %s", phi.FunctionFullPath(actual), m.ref)
}

func (m *beFunctionMatcher) NegatedFailureMessage(actual any) string {
	return fmt.Sprintf("expected %s to be different of %s", phi.FunctionFullPath(actual), m.ref)
}

func BeFunction(f any) types.GomegaMatcher {
	if !phi.IsFunction(f) {
		panic(fmt.Errorf("input parameter must be a function, received `%T`", f))
	}
	return &beFunctionMatcher{ref: phi.FunctionFullPath(f)}
}
