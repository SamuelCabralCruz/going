package result

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple"
)

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value:   value,
		isError: false,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		err:     err,
		isError: true,
	}
}

func Errorf[T any](format string, a ...any) Result[T] {
	return Error[T](fmt.Errorf(format, a...))
}

func FromTuple[T any](value T, err error) Result[T] {
	if err != nil {
		return Error[T](err)
	}
	return Ok(value)
}

func FromSupplier[T any](supply fn.Supplier[T]) Result[T] {
	return FromTuple(fn.Safe(supply))
}

func FromProducer[T any](produce fn.Producer[T]) Result[T] {
	return FromTuple(fn.Try(produce))
}

type Result[T any] struct {
	value   T
	err     error
	isError bool
}

func (r Result[T]) IsOk() bool {
	return !r.isError
}

func (r Result[T]) IsError() bool {
	return r.isError
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

func (r Result[T]) GetOrPanic() T {
	return tuple.GetOrPanic(r.Get())
}

func (r Result[T]) OrEmpty() T {
	return r.value
}

func (r Result[T]) OrElseGet(supply fn.Supplier[T]) T {
	if r.isError {
		return supply()
	}
	return r.value
}

func (r Result[T]) OrElse(value T) T {
	return r.OrElseGet(func() T { return value })
}
