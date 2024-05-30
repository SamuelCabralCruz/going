package validation

import (
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/phi"
)

func IgnoreOk[T any](value T, _ bool) T {
	return value
}

func IgnoreValue[T any](_ T, ok bool) bool {
	return ok
}

func GetOrEmpty[T any](value T, ok bool) T {
	if !ok {
		return phi.Empty[T]()
	}
	return value
}

func GetOrPanic[T any](value T, ok bool) T {
	_, err := ToAssertion(value, ok)
	if err != nil {
		panic(err)
	}
	return value
}

func PanicIfNotOk[T any](value T, ok bool) {
	GetOrPanic(value, ok)
}

func Ok[T any](value T) (T, bool) {
	return value, true
}

func NotOk[T any]() (T, bool) {
	return phi.Empty[T](), false
}

func Reverse[T any](value T, ok bool) (bool, T) {
	return ok, value
}

func FromReversed[T any](ok bool, value T) (T, bool) {
	return value, ok
}

func ToAssertion[T any](value T, ok bool) (T, error) {
	if !ok {
		return phi.Empty[T](), newInvalidValueError(value)
	}
	return value, nil
}

func Switch[T any, U any](value T, ok bool) func(mapper typing.Transformer[T, U], supplier typing.Supplier[U]) U {
	return func(transformer typing.Transformer[T, U], supplier typing.Supplier[U]) U {
		if !ok {
			return supplier()
		}
		return transformer(value)
	}
}
