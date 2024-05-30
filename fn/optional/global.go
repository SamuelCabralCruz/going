package optional

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/samber/lo"
)

func Transform[T any, U any](value Optional[T], transformer typing.Transformer[T, U]) Optional[U] {
	if value.IsPresent() {
		return FromAssertion(fn.SafeTransformer(transformer, value.GetOrPanic()))
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

func Combine[T any](opts ...Optional[T]) Optional[[]T] {
	return OfNullable(FilterPresent[T](opts...))
}
