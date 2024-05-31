package xpctd

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/phi"
)

func Computed[T any](describe func(actual T) string) Expected[T] {
	r := newReporter[T]()
	r.expected = describe
	return r
}

func Formatted[T any](format string, a ...any) Expected[T] {
	return Computed(func(actual T) string {
		return fmt.Sprintf(format, a...)
	})
}

func The[T any](description string) Expected[T] {
	return Formatted[T]("the %s", description)
}

func Type[T any]() Expected[T] {
	return Formatted[T](phi.BaseTypeName[T]())
}

func Value[T any]() Expected[T] {
	return Formatted[T]("value")
}

func Actual[T any]() Expected[T] {
	return Computed[T](func(actual T) string {
		return fmt.Sprintf(`"%+v"`, actual)
	})
}
