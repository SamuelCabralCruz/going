package fn

import (
	"github.com/SamuelCabralCruz/went/roar"
)

func Identity[T any](value T) T {
	return value
}

func Try[T any](produce Producer[T]) (value T, err error) {
	rec := func() {
		if r := recover(); r != nil {
			err = roar.AsError(r)
		}
	}
	defer rec()
	value, err = produce()
	return
}

func Safe[T any](supply Supplier[T]) (value T, err error) {
	return Try(func() (T, error) {
		return supply(), nil
	})
}

func Prevent(callable Callable) error {
	_, err := Safe(func() string {
		callable()
		return "ok"
	})
	return err
}
