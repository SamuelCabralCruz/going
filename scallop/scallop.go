package scallop

import (
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
)

func AsAnySlice[T any](values []T) []any {
	return lo.Map(values, func(value T, _ int) any {
		return any(value)
	})
}

func Copy[T any](values []T) []T {
	return append([]T{}, values...)
}

func Pop[T any](values []T) (T, []T, error) {
	if len(values) == 0 {
		return phi.Empty[T](), phi.Empty[[]T](), newIndexOutOfBoundsError(0, 0)
	}
	return values[0], values[1:], nil
}
