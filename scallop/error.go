package scallop

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/roar"
)

type IndexOutOfBoundsError struct {
	roar.Roar[IndexOutOfBoundsError]
}

func newIndexOutOfBoundsError(index int, length int) IndexOutOfBoundsError {
	return IndexOutOfBoundsError{
		Roar: roar.New[IndexOutOfBoundsError](
			"access to non existing element has been intercepted",
			roar.WithField("index", index),
			roar.WithField("bounds", fmt.Sprintf("[0,%d]", length))),
	}
}
