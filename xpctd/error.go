package xpctd

import (
	"github.com/SamuelCabralCruz/going/roar"
)

type ExpectationError struct {
	roar.Roar[ExpectationError]
}

func newExpectationError(message string) ExpectationError {
	return ExpectationError{
		roar.New[ExpectationError](message),
	}
}
