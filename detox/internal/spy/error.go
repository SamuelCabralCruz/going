package spy

import (
	"github.com/SamuelCabralCruz/went/roar"
)

type InvalidCallIndexError struct {
	roar.Roar[InvalidCallIndexError]
}

func newInvalidCallIndexError(mockName string, methodName string, received int, nbCalls int) InvalidCallIndexError {
	return InvalidCallIndexError{
		roar.New[InvalidCallIndexError](
			"invalid call index provided",
			roar.WithField("mock", mockName),
			roar.WithField("method", methodName),
			roar.WithField("received", received),
			roar.WithField("nbCalls", nbCalls))}
}
