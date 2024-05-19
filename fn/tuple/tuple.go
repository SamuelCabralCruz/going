package tuple

import (
	"github.com/SamuelCabralCruz/went/phi"
)

func Ignore[T any](value T, _ error) T {
	return value
}

func GetOrPanic[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func FromValue[T any](value T) (T, error) {
	return value, nil
}

func FromError[T any](err error) (T, error) {
	return phi.Empty[T](), err
}

func Swap[T any, U any](left T, right U) (U, T) {
	return right, left
}
