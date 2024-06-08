package phi

import (
	"fmt"
	"reflect"
)

func IsTypeOf[T any](value any) (T, bool) {
	v, ok := value.(T)
	return v, ok
}

func IsZero(value any) bool {
	if value == nil {
		return true
	}
	if IsMap(value) {
		return compareBySerialization(value, map[any]any{})
	}
	if IsSlice(value) {
		return compareBySerialization(value, []any{})
	}
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func IsNotZero(value any) bool {
	return !IsZero(value)
}

func compareBySerialization(value any, zeroValue any) bool {
	return fmt.Sprintf("%+v", value) == fmt.Sprintf("%+v", zeroValue)
}

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
