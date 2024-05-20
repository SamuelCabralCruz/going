package result

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/samber/lo"
)

func Transform[T any, U any](value Result[T], transform fn.Transformer[T, U]) Result[U] {
	if value.IsOk() {
		return Ok(transform(value.GetOrPanic()))
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
