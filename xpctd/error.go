package xpctd

import (
	"github.com/SamuelCabralCruz/went/roar"
)

type ExpectationError struct {
	roar.Roar[ExpectationError]
}

func newExpectationError(message string) ExpectationError {
	return ExpectationError{
		roar.New[ExpectationError](message),
	}
}
