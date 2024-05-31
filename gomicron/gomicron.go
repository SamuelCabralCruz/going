package gomicron

import "github.com/SamuelCabralCruz/went/gomicron/reporter"

type Matcher[T any] func(T) (bool, error)

type MatcherDefinition[T any] struct {
	Matcher  Matcher[T]
	Reporter reporter.Reporter[T]
}
