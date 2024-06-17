package matcher

import (
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/gomicron"
	"github.com/SamuelCabralCruz/going/xpctd"
	"github.com/onsi/gomega/types"
)

func HaveBeenCalled() types.GomegaMatcher {
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[detox.Assertable]{
		Matcher: func(actual detox.Assertable) (bool, error) {
			return actual.Assert().HasBeenCalled(), nil
		},
		Reporter: xpctd.Computed[detox.Assertable](
			func(actual detox.Assertable) string {
				return actual.Describe()
			}).
			ToHaveFormatted("been called at least once").
			ButWasFormatted("not called"),
	})
}
