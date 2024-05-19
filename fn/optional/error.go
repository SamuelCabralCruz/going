package optional

import "github.com/SamuelCabralCruz/went/roar"

type NoSuchElementError struct {
	roar.Roar[NoSuchElementError]
}

func newNoSuchElementError() NoSuchElementError {
	return NoSuchElementError{
		roar.New[NoSuchElementError]("access to value of empty optional has been intercepted")}
}
