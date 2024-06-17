package matcher

import (
	"github.com/SamuelCabralCruz/going/fn/result"
	"github.com/SamuelCabralCruz/going/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/going/gomicron"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/SamuelCabralCruz/going/trust"
	"github.com/SamuelCabralCruz/going/xpctd"
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
