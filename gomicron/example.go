package gomicron

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/onsi/gomega/types"
)

func BeFunction(f uintptr) types.GomegaMatcher {
	if !phi.IsFunction(f) {
		panic(fmt.Errorf("input parameter must be a function, received `%T`", f))
	}
	ref := phi.FunctionFullPath(f)
	return ToGomegaMatcher(MatcherDefinition[any]{
		Matcher: func(actual any) (bool, error) {
			if !phi.IsFunction(actual) {
				return false, fmt.Errorf("actual must be a function, received `%T`", actual)
			}
			return ref == phi.FunctionFullPath(actual), nil
		},
		Reporter: func(actual any, polarity Polarity) string {
			return "some reporting message"
		},
	})
}

func BeAStringOfLength(length int) types.GomegaMatcher {
	return ToGomegaMatcher(MatcherDefinition[string]{
		Matcher: func(actual string) (bool, error) {
			return len(actual) == length, nil
		},
		Reporter: func(actual string, polarity Polarity) string {
			return "some reporting message"
		},
	})
}
