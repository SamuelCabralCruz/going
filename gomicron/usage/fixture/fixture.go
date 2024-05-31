//go:build test

package fixture

import (
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/xpctd"
	"github.com/onsi/gomega/types"
)

func BeSomeCustomMatcher() types.GomegaMatcher {
	return gomicron.ToGomegaMatcher(gomicron.MatcherDefinition[string]{
		Matcher:  func(actual string) (bool, error) { return len(actual) == 6, nil },
		Reporter: xpctd.Actual[string](),
	})
}
