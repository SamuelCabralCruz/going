package fn

import "github.com/samber/lo"

func ChainMappers[T any](mappers ...Mapper[T]) Mapper[T] {
	return lo.Reduce(mappers[1:], func(agg Mapper[T], mapper Mapper[T], _ int) Mapper[T] {
		return func(value T) T {
			return mapper(agg(value))
		}
	}, mappers[0])
}

func ChainTransformers[T any, U any, V any](t1 Transformer[T, U], t2 Transformer[U, V]) Transformer[T, V] {
	return func(value T) V {
		return t2(t1(value))
	}
}
