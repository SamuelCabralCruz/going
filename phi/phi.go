package phi

import (
	"fmt"
	"reflect"
)

func Empty[T any]() (t T) {
	return
}

func Kind[T any]() reflect.Kind {
	return reflect.ValueOf(Empty[T]()).Kind()
}

func Type[T any]() string {
	empty := Empty[T]()
	typeOf := reflect.TypeOf(empty)
	if typeOf == nil {
		t := (*T)(nil)
		return reflect.TypeOf(t).Elem().Name()
	}
	return typeOf.Name()
}

func PkgPath[T any]() string {
	empty := Empty[T]()
	typeOf := reflect.TypeOf(empty)
	if typeOf == nil {
		t := (*T)(nil)
		return reflect.TypeOf(t).Elem().PkgPath()
	}
	return typeOf.PkgPath()
}

func UniqueIdentifier[T any]() string {
	return fmt.Sprintf("%s.%s[%s]", PkgPath[T](), Type[T](), Kind[T]().String())
}

func IsImplementing[T any](t any) bool {
	_, ok := t.(T)
	return ok
}
