package validation

import (
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
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
	return assertion.GetOrPanic(ToAssertion(value, ok))
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

func FromReversed[T any](ok bool, value T) (T, bool) {
	return value, ok
}

func ToAssertion[T any](value T, ok bool) (T, error) {
	return ToAssertionWithError[T](value, ok)(newInvalidValueError(value))
}

func ToAssertionWithError[T any](value T, ok bool) func(error) (T, error) {
	return func(err error) (T, error) {
		if !ok {
			return phi.Empty[T](), err
		}
		return value, nil
	}
}

func Switch[T any, U any](value T, ok bool) func(typing.Transformer[T, U], typing.Supplier[U]) U {
	return func(onOk typing.Transformer[T, U], onNotOk typing.Supplier[U]) U {
		if !ok {
			return onNotOk()
		}
		return onOk(value)
	}
}
