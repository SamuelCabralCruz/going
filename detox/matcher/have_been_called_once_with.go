package matcher

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/xpctd"
	"github.com/onsi/gomega/types"
	"strings"
)

func HaveBeenCalledOnceWith(args ...any) types.GomegaMatcher {
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[detox.Assertable]{
		Matcher: func(actual detox.Assertable) (bool, error) {
			return actual.Assert().HasBeenCalledOnceWith(args...), nil
		},
		Reporter: xpctd.Computed[detox.Assertable](
			func(actual detox.Assertable) string {
				return actual.Describe()
			}).
			ToHaveFormatted("been called once with following args:\n\t\t%s\n", describeArgs(args)).
			ButReceived(func(actual detox.Assertable) string {
				return fmt.Sprintf("calls were:\n\t\t%s", strings.Join(describeCalls(actual.Calls()), "\n\t\t"))
			}),
	})
}
