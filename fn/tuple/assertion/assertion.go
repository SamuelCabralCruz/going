package assertion

import (
	"github.com/SamuelCabralCruz/going/fn/typing"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/SamuelCabralCruz/going/roar"
)

func IgnoreError[T any](value T, _ error) T {
	return value
}

func IgnoreValue[T any](_ T, err error) error {
	return err
}

func GetOrEmpty[T any](value T, err error) T {
	if err != nil {
		return phi.Empty[T]()
	}
	return value
}

func GetOrPanic[T any](value T, err error) T {
	roar.PanicIfError(err)
	return value
}

func PanicIfError[T any](value T, err error) {
	GetOrPanic(value, err)
}

func FromValue[T any](value T) (T, error) {
	return value, nil
}

func FromError[T any](err error) (T, error) {
	return phi.Empty[T](), err
}

func FromReversed[T any](err error, value T) (T, error) {
	return value, err
}

func ToValidation[T any](value T, err error) (T, bool) {
	if err != nil {
		return phi.Empty[T](), false
	}
	return value, true
}

func Switch[T any, U any](
	value T,
	err error,
) func(typing.Transformer[T, U], typing.Transformer[error, U]) U {
	return func(onOk typing.Transformer[T, U], onError typing.Transformer[error, U]) U {
		if err != nil {
			return onError(err)
		}
		return onOk(value)
	}
}
