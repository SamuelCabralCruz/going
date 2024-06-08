package phi

import (
	"fmt"
	"reflect"
)

func InterfaceToPtr[T any]() *T {
	return (*T)(nil)
}

func InterfaceToType[T any]() reflect.Type {
	return reflect.TypeOf(InterfaceToPtr[T]()).Elem()
}

func GetInterfaceMethodByName[T any](methodName string) (reflect.Method, bool) {
	return InterfaceToType[T]().MethodByName(methodName)
}

func GetMatchingInterfaceMethod[T any, U any](method U) (reflect.Method, error) {
	methodName := FunctionName(method)
	if v, ok := GetInterfaceMethodByName[T](methodName); ok {
		expectedMethodType := v.Type
		observedMethodType := Value(method).Type()
		if expectedMethodType == observedMethodType {
			return v, nil
		}
		return reflect.Method{}, fmt.Errorf(
			`interface method type "%s" does not match provided method type "%s"`,
			expectedMethodType,
			observedMethodType)
	}
	return reflect.Method{}, fmt.Errorf(
		`interface "%s" does not have a method named "%s"`,
		BaseTypeName[T](), methodName)
}
