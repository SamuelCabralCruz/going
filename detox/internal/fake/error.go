package fake

import (
	"github.com/SamuelCabralCruz/went/roar"
)

type MissingFakeImplementationError struct {
	roar.Roar[MissingFakeImplementationError]
}

func newMissingFakeImplementationError(mockName string, methodName string) MissingFakeImplementationError {
	return MissingFakeImplementationError{
		roar.New[MissingFakeImplementationError](
			"no fake has been registered",
			roar.WithField("mock", mockName),
			roar.WithField("method", methodName))}
}
