package phi

import (
	"fmt"
	"reflect"
)

func Empty[T any]() (t T) {
	return
}

func InterfaceToPtr[T any]() *T {
	return (*T)(nil)
}

func Value[T any]() reflect.Value {
	return reflect.ValueOf(Empty[T]())
}

func Type[T any]() reflect.Type {
	empty := Empty[T]()
	typeOf := reflect.TypeOf(empty)
	if typeOf == nil {
		return reflect.TypeOf(InterfaceToPtr[T]()).Elem()
	}
	return typeOf
}

func Kind[T any]() reflect.Kind {
	return Value[T]().Kind()
}

func PkgPath[T any]() string {
	return Type[T]().PkgPath()
}

func UniqueIdentifier[T any]() string {
	return fmt.Sprintf("%s.%s[%s]", PkgPath[T](), Type[T](), Kind[T]())
}

func IsImplementing[T any](t any) bool {
	_, ok := t.(T)
	return ok
}
