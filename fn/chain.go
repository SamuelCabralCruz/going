package fn

import (
	"github.com/SamuelCabralCruz/going/fn/typing"
	"github.com/samber/lo"
)

func ChainMappers[T any](mappers ...typing.Mapper[T]) typing.Mapper[T] {
	return lo.Reduce(mappers[1:], func(agg typing.Mapper[T], mapper typing.Mapper[T], _ int) typing.Mapper[T] {
		return func(value T) T {
			return mapper(agg(value))
		}
	}, mappers[0])
}

func ChainTransformers[T any, U any, V any](t1 typing.Transformer[T, U], t2 typing.Transformer[U, V]) typing.Transformer[T, V] {
	return func(value T) V {
		return t2(t1(value))
	}
}
