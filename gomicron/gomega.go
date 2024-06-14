package gomicron

import (
	"github.com/SamuelCabralCruz/went/xpctd/construct"
	"github.com/onsi/gomega/types"
)

func ToGomegaMatcher[T any](definition MatcherDefinition[T]) types.GomegaMatcher {
	return customGomegaMatcher[T]{definition}
}

type customGomegaMatcher[T any] struct {
	definition MatcherDefinition[T]
}

var _ types.GomegaMatcher = customGomegaMatcher[any]{}

func (m customGomegaMatcher[T]) Match(actual any) (bool, error) {
	if v, ok := actual.(T); ok {
		return m.definition.Matcher(v)
	}
	return false, construct.TypeMismatchReporter[T]().Error(actual)
}

func (m customGomegaMatcher[T]) FailureMessage(actual any) string {
	return m.definition.Reporter.Report(actual.(T))
}

func (m customGomegaMatcher[T]) NegatedFailureMessage(actual any) string {
	return m.definition.Reporter.Negative().Report(actual.(T))
}
