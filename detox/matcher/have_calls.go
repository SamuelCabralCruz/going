package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/gomicron"
	"github.com/SamuelCabralCruz/going/xpctd"
	"github.com/onsi/gomega/types"
)

func HaveCalls(calls ...detox.Call) types.GomegaMatcher {
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[detox.Assertable]{
		Matcher: func(actual detox.Assertable) (bool, error) {
			return actual.Assert().HasCalls(calls...), nil
		},
		Reporter: xpctd.Computed[detox.Assertable](
			func(actual detox.Assertable) string {
				return actual.Describe()
			}).
			ToHaveFormatted("following calls in any order:\n%s\n", describeCalls(toCommonCalls(calls))).
			ButReceived(func(actual detox.Assertable) string {
				return fmt.Sprintf("calls were:\n%s", describeCalls(actual.Calls()))
			}),
	})
}
