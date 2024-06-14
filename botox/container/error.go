package container

import (
	"github.com/SamuelCabralCruz/went/roar"
	"reflect"
)

type NoCandidateFoundError struct {
	roar.Roar[NoCandidateFoundError]
}

func newNoCandidateFoundError(typeof reflect.Type) NoCandidateFoundError {
	return NoCandidateFoundError{
		roar.New[NoCandidateFoundError](
			"no provider have been registered for requested type",
			roar.WithField("type", typeof)),
	}
}

type TooManyCandidatesFoundError struct {
	roar.Roar[TooManyCandidatesFoundError]
}

func newTooManyCandidatesFoundError(typeof reflect.Type, n int) TooManyCandidatesFoundError {
	return TooManyCandidatesFoundError{
		roar.New[TooManyCandidatesFoundError](
			"too many providers have been registered for requested type",
			roar.WithField("type", typeof),
			roar.WithField("expected", 1),
			roar.WithField("received", n)),
	}
}
