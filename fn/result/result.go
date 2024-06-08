package result

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/fn/tuple/validation"
	"github.com/SamuelCabralCruz/went/fn/typing"
)

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: value,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func Errorf[T any](format string, a ...any) Result[T] {
	return Error[T](fmt.Errorf(format, a...))
}

func FromAssertion[T any](value T, err error) Result[T] {
	if err != nil {
		return Error[T](err)
	}
	return Ok(value)
}

func FromValidation[T any](value T, ok bool) Result[T] {
	return FromAssertion(validation.ToAssertion(value, ok))
}

func FromSupplier[T any](supplier typing.Supplier[T]) Result[T] {
	return FromAssertion(fn.SafeSupplier(supplier))
}

func FromProducer[T any](producer typing.Producer[T]) Result[T] {
	return FromAssertion(fn.SafeProducer(producer))
}

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsError() bool {
	return !r.IsOk()
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

func (r Result[T]) GetOrPanic() T {
	return assertion.GetOrPanic(r.Get())
}

func (r Result[T]) GetOrPanicWith(err error) T {
	if r.IsError() {
		panic(err)
	}
	return r.value
}

func (r Result[T]) OrEmpty() T {
	return r.value
}

func (r Result[T]) OrElseGet(supplier typing.Supplier[T]) T {
	if r.IsError() {
		return FromSupplier(supplier).OrEmpty()
	}
	return r.value
}

func (r Result[T]) OrElse(value T) T {
	if r.IsError() {
		return value
	}
	return r.value
}

func (r Result[T]) FlatMap(mapper typing.Mapper[Result[T]]) Result[T] {
	return assertion.Switch[Result[T], Result[T]](fn.SafeMapper(mapper, r))(fn.Identity[Result[T]], Error[T])
}

func (r Result[T]) Map(mapper typing.Mapper[T]) Result[T] {
	if r.IsOk() {
		return assertion.Switch[T, Result[T]](fn.SafeMapper(mapper, r.value))(Ok[T], Error[T])
	}
	return r
}

func (r Result[T]) MapError(transformer typing.Transformer[error, T]) Result[T] {
	if r.IsError() {
		return assertion.Switch[T, Result[T]](fn.SafeTransformer(transformer, r.err))(Ok[T], Error[T])
	}
	return r
}

func (r Result[T]) SwitchMap(onValue typing.Mapper[T], onError typing.Transformer[error, T]) Result[T] {
	if r.IsOk() {
		return r.Map(onValue)
	}
	return r.MapError(onError)
}

func (r Result[T]) IfOk(consume typing.Consumer[T]) {
	if r.IsOk() {
		consume(r.value)
	}
}

func (r Result[T]) IfError(consume typing.Consumer[error]) {
	if r.IsError() {
		consume(r.err)
	}
}

func (r Result[T]) Switch(onOk typing.Consumer[T], onError typing.Consumer[error]) {
	r.IfOk(onOk)
	r.IfError(onError)
}

func (r Result[T]) Filter(predicate typing.Predicate[T]) Result[T] {
	if r.IsOk() {
		filterWithPredicate := func(predicated bool) Result[T] {
			if predicated {
				return r
			}
			return Error[T](newFilteredValueError(r.value))
		}
		assertion.Switch[bool, Result[T]](fn.SafePredicate(predicate, r.value))(filterWithPredicate, Error[T])
	}
	return r
}
