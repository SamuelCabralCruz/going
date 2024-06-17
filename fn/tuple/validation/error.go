package validation

import "github.com/SamuelCabralCruz/going/roar"

type InvalidValueError struct {
	roar.Roar[InvalidValueError]
}

func newInvalidValueError(value any) InvalidValueError {
	return InvalidValueError{
		Roar: roar.New[InvalidValueError](
			"value has failed its validation process",
			roar.WithField("value", value)),
	}
}
