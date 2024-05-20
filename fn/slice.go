package fn

import "github.com/samber/lo"

func AsAnySlice[T any](values []T) []any {
	return lo.Map(values, func(value T, _ int) any {
		return any(value)
	})
}
