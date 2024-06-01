package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/xpctd"
	"github.com/onsi/gomega/types"
	"strings"
)

func HaveBeenCalledTimesWith(times int, args ...any) types.GomegaMatcher {
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[detox.Assertable]{
		Matcher: func(actual detox.Assertable) (bool, error) {
			return actual.Assert().HasBeenCalledTimesWith(times, args...), nil
		},
		Reporter: xpctd.Computed[detox.Assertable](
			func(actual detox.Assertable) string {
				return actual.Describe()
			}).
			ToHaveFormatted("been called %d times with following args:\n\t\t%s\n", times, describeArgs(args...)).
			ButReceived(func(actual detox.Assertable) string {
				return fmt.Sprintf("calls were:\n\t\t%s", strings.Join(describeCalls(actual), "\n\t\t"))
			}),
	})
}
