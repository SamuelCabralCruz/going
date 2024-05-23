package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/onsi/gomega/types"
)

type haveBeenCalledMatcher struct{}

var _ types.GomegaMatcher = &haveBeenCalledMatcher{}

func (m *haveBeenCalledMatcher) Match(actual any) (bool, error) {
	if v, ok := actual.(detox.Assertable); ok {
		return v.Assert().HasBeenCalled(), nil
	}
	return false, fmt.Errorf("actual `%+v` is not a mocked method", actual)
}

func (m *haveBeenCalledMatcher) FailureMessage(actual any) string {
	mock := actual.(detox.Assertable)
	return fmt.Sprintf("expected %s to have been called", mock.Describe())
}

func (m *haveBeenCalledMatcher) NegatedFailureMessage(actual any) string {
	mock := actual.(detox.Assertable)
	return fmt.Sprintf("expected %s not to have been called", mock.Describe())
}

func HaveBeenCalled() types.GomegaMatcher {
	return &haveBeenCalledMatcher{}
}
