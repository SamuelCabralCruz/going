package fn

import "github.com/samber/lo"

func AsAnySlice[T any](values []T) []any {
	return lo.Map(values, func(value T, _ int) any {
		return any(value)
	})
}

func Copy[T any](values []T) []T {
	return append([]T{}, values...)
}

func Pop[T any](values []T) (T, []T, error) {
	// TODO: complete with roar
	// TODO: https://go.dev/play/p/Pdzc4bhMhIE
	return values[0], values[1:], nil
}
