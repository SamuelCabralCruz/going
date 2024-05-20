package phi

import (
	"fmt"
	"reflect"
)

func IsFunction(value any) bool {
	return IsKind(value, reflect.Func)
}

func IsMap(value any) bool {
	return IsKind(value, reflect.Map)
}

func IsSlice(value any) bool {
	return IsKind(value, reflect.Slice)
}

func IsKind(value any, kind reflect.Kind) bool {
	return reflect.TypeOf(value).Kind() == kind
}

func IsZero(value any) bool {
	if IsMap(value) {
		return compareBySerialization(value, map[any]any{})
	}
	if IsSlice(value) {
		return compareBySerialization(value, []any{})
	}
	return value == nil || reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func compareBySerialization(value any, zeroValue any) bool {
	return fmt.Sprintf("%+v", value) == fmt.Sprintf("%+v", zeroValue)
}

func IsImplementing[T any](value any) bool {
	_, ok := value.(T)
	return ok
}
