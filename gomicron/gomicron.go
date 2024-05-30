package gomicron

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/onsi/gomega/types"
)

type MatcherDefinition[T any] struct {
	Matcher  func(T) (bool, error)
	Reporter func(T, Polarity) string
}

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
	return false, fmt.Errorf(`expected "%+v" to be a "%s" but was a "%T"`, actual, phi.TypeName[T](), actual)
}

func (m customGomegaMatcher[T]) FailureMessage(actual any) string {
	return m.definition.Reporter(actual.(T), Positive)
}

func (m customGomegaMatcher[T]) NegatedFailureMessage(actual any) string {
	return m.definition.Reporter(actual.(T), Negative)
}
