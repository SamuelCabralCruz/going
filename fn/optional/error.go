package optional

import "github.com/SamuelCabralCruz/went/roar"

type NoSuchElementError struct {
	roar.Roar[NoSuchElementError]
}

func newNoSuchElementError() NoSuchElementError {
	return NoSuchElementError{
		Roar: roar.New[NoSuchElementError]("access to a missing element has been intercepted")}
}
