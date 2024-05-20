package optional

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/samber/lo"
)

func Transform[T any, U any](value Optional[T], transform fn.Transformer[T, U]) Optional[U] {
	if value.IsPresent() {
		return Of(transform(value.GetOrPanic()))
	}
	return Empty[U]()
}

func FilterPresent[T any](values ...Optional[T]) []T {
	return lo.Reduce(values, func(agg []T, value Optional[T], _ int) []T {
		if value.IsPresent() {
			agg = append(agg, value.GetOrPanic())
		}
		return agg
	}, []T{})
}
