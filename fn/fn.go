package fn

import (
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
	"github.com/samber/mo"
)

func Accumulate[T any](results ...mo.Result[T]) mo.Result[[]T] {
	var acc []T
	var errors []error
	lo.ForEach(results,
		func(result mo.Result[T], _ int) {
			if result.IsOk() {
				acc = append(acc, result.MustGet())
			} else {
				errors = append(errors, result.Error())
			}
		})
	if len(errors) > 0 {
		return mo.Err[[]T](NewAggregatedError(errors...))
	}
	return mo.Ok(acc)
}

func ErrorHasTuple[T any](err error) (T, error) {
	empty := phi.EmptyOf[T]()
	return empty, err
}

func ValueHasTuple[T any](value T) (T, error) {
	return value, nil
}

func AsResultable[T any](f func() T) func() mo.Result[T] {
	return func() mo.Result[T] {
		return mo.Ok(f())
	}
}
