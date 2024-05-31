package matcher

import (
	"github.com/SamuelCabralCruz/went/fn/result"
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/SamuelCabralCruz/went/trust"
	"github.com/SamuelCabralCruz/went/xpctd"
	"github.com/onsi/gomega/types"
)

func BeFunction(f any) types.GomegaMatcher {
	assertion.PanicIfError(trust.AssertIsFunction(f))
	ref := phi.FunctionFullPath(f)
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[any]{
		Matcher: func(actual any) (bool, error) {
			return result.Transform(result.FromAssertion(
				trust.AssertIsFunction(actual)),
				func(_ any) bool {
					return ref == phi.FunctionFullPath(actual)
				}).Get()
		},
		Reporter: xpctd.Computed[any](
			func(actual any) string {
				return phi.FunctionFullPath(actual)
			}).
			ToBeFormatted("identical to %s", ref),
	})
}
