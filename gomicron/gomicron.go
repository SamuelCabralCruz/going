package gomicron

import (
	"github.com/SamuelCabralCruz/going/xpctd"
)

type Matcher[T any] func(T) (bool, error)

type MatcherDefinition[T any] struct {
	Matcher  Matcher[T]
	Reporter xpctd.Reporter[T]
}
