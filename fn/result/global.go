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

func Combine[T any](results ...Result[T]) Result[[]T] {
	errors := FilterError(results...)
	if len(errors) > 0 {
		return Error[[]T](roar.Aggregate(errors...))
	}
	return Ok(FilterOk(results...))
}
