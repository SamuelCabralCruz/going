package result

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/roar"
	"github.com/samber/lo"
)

func Transform[T any, U any](value Result[T], transformer typing.Transformer[T, U]) Result[U] {
	if value.IsOk() {
		return FromAssertion(fn.SafeTransformer(transformer, value.GetOrPanic()))
	}
	return Error[U](value.Error())
}

func FilterOk[T any](values ...Result[T]) []T {
	return lo.Reduce(values, func(agg []T, value Result[T], _ int) []T {
		if value.IsOk() {
			agg = append(agg, value.GetOrPanic())
		}
		return agg
	}, []T{})
}

func FilterError[T any](values ...Result[T]) []error {
	return lo.Reduce(values, func(agg []error, value Result[T], _ int) []error {
		if value.IsError() {
			agg = append(agg, value.Error())
		}
		return agg
	}, []error{})
}

// TODO: validate contract below
func Combine[T any](results ...Result[T]) Result[[]T] {
	var acc []T
	var errors []error
	lo.ForEach(results,
		func(result Result[T], _ int) {
			if result.IsOk() {
				acc = append(acc, result.GetOrPanic())
			} else {
				errors = append(errors, result.Error())
			}
		})
	if len(errors) == 1 {
		return Error[[]T](errors[0])
	}
	if len(errors) > 1 {
		return Error[[]T](roar.NewAggregatedError(errors...))
	}
	return Ok(acc)
}
