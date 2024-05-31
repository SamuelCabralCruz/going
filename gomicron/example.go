package gomicron

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/gomicron/reporter"
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
		Reporter: reporter.Computed[any](
			func(actual any) string {
				return phi.FunctionFullPath(actual)
			}).
			ToBeFormatted("identical to %s", ref),
	})
}

func BeAStringOfLength(length int) types.GomegaMatcher {
	return ToGomegaMatcher(MatcherDefinition[string]{
		Matcher: func(actual string) (bool, error) {
			return len(actual) == length, nil
		},
		Reporter: reporter.Actual[string]().
			ToHaveFormatted(`length of "%d"`, length).
			ButHad(func(actual string) string {
				return fmt.Sprintf(`length of "%d"`, len(actual))
			}),
	})
}
