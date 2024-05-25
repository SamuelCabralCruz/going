package phi

import (
	"fmt"
	"reflect"
)

func Empty[T any]() (t T) {
	return
}

func Value[T any](value T) reflect.Value {
	return reflect.ValueOf(value)
}

func Type[T any]() reflect.Type {
	empty := Empty[T]()
	typeOf := reflect.TypeOf(empty)
	if typeOf == nil {
		return InterfaceToType[T]()
	}
	return typeOf
}

func Kind[T any]() reflect.Kind {
	return Type[T]().Kind()
}

func PkgPath[T any]() string {
	return Type[T]().PkgPath()
}

func UniqueIdentifier[T any]() string {
	return fmt.Sprintf("%s.%s[%s]", PkgPath[T](), Type[T](), Kind[T]())
}
