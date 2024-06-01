package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/xpctd"
	"github.com/onsi/gomega/types"
	"strings"
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
			ToHaveFormatted("following calls:\n\t\t%s\n", strings.Join(describeCalls(toCommonCalls(calls)), "\n\t\t")).
			ButReceived(func(actual detox.Assertable) string {
				return fmt.Sprintf("calls were:\n\t\t%s", strings.Join(describeCalls(actual.Calls()), "\n\t\t"))
			}),
	})
}
