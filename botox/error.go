package botox

import (
	"github.com/SamuelCabralCruz/went/roar"
)

type ProvidingLoopError struct {
	roar.Roar[ProvidingLoopError]
}

func NewProvidingLoopError() ProvidingLoopError {
	return ProvidingLoopError{
		roar.New[ProvidingLoopError]("unexpected error occurred during resolution loop"),
	}
}

type NoCandidateFoundError struct {
	roar.Roar[NoCandidateFoundError]
}

func NewNoCandidateFoundError(typeof string) NoCandidateFoundError {
	return NoCandidateFoundError{
		roar.New[NoCandidateFoundError](
			"no provider have been registered for requested type",
			roar.WithField("type", typeof)),
	}
}

type TooManyCandidatesFoundError struct {
	roar.Roar[TooManyCandidatesFoundError]
}

func NewTooManyCandidatesFoundError(typeof string, n int) TooManyCandidatesFoundError {
	return TooManyCandidatesFoundError{
		roar.New[TooManyCandidatesFoundError](
			"too many providers have been registered for requested type",
			roar.WithField("type", typeof),
			roar.WithField("expected", 1),
			roar.WithField("received", n)),
	}
}
