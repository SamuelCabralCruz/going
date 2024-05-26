package result

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple"
	"github.com/SamuelCabralCruz/went/roar"
	"github.com/samber/lo"
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

func (r Result[T]) GetOrPanicWith(err error) T {
	if r.IsError() {
		panic(err)
	}
	return r.value
}

func (r Result[T]) OrEmpty() T {
	return r.value
}

func (r Result[T]) OrElseTry(produce fn.Producer[T]) Result[T] {
	if r.IsError() {
		return FromProducer(produce)
	}
	return Ok(r.value)
}

func (r Result[T]) OrElseGet(supply fn.Supplier[T]) T {
	return r.OrElseTry(fn.ToProducer(supply)).OrEmpty()
}

func (r Result[T]) OrElse(value T) T {
	return r.OrElseGet(fn.ToSupplier(value))
}

func (r Result[T]) TryFlatMap(mapper fn.TryableMapper[Result[T]]) Result[T] {
	if r.IsOk() {
		maybe, err := mapper(r)
		if err != nil {
			return Error[T](err)
		}
		return maybe
	}
	return r
}

func (r Result[T]) TryFlatMapError(mapper fn.TryableMapper[Result[T]]) Result[T] {
	if r.IsError() {
		maybe, err := mapper(r)
		if err != nil {
			return Error[T](err)
		}
		return maybe
	}
	return r
}

func (r Result[T]) FlatMap(mapper fn.Mapper[Result[T]]) Result[T] {
	return r.TryFlatMap(fn.ToTryableMapper(mapper))
}

func (r Result[T]) FlatMapError(mapper fn.Mapper[Result[T]]) Result[T] {
	return r.TryFlatMapError(fn.ToTryableMapper(mapper))
}

func (r Result[T]) TryMap(mapper fn.TryableMapper[T]) Result[T] {
	if r.IsOk() {
		return FromTuple[T](mapper(r.value))
	}
	return r
}

func (r Result[T]) TryMapError(transform fn.TryableTransformer[error, T]) Result[T] {
	if r.IsError() {
		return FromTuple[T](transform(r.err))
	}
	return r
}

func (r Result[T]) Map(mapper fn.Mapper[T]) Result[T] {
	return r.TryMap(fn.ToTryableMapper(mapper))
}

func (r Result[T]) MapError(transform fn.Transformer[error, T]) Result[T] {
	return r.TryMapError(fn.ToTryableTransformer(transform))
}

func (r Result[T]) IfOk(consume fn.Consumer[T]) {
	if r.IsOk() {
		consume(r.value)
	}
}

func (r Result[T]) IfError(consume fn.Consumer[error]) {
	if r.IsError() {
		consume(r.err)
	}
}

// TODO: review this contract
func (r Result[T]) TrySwitchMap(valueMapper fn.TryableMapper[T], errorMapper fn.TryableErrorMapper[T]) Result[T] {
	if r.IsError() {
		return FromTuple[T](errorMapper(r.err))
	}
	return FromTuple(valueMapper(r.value))
}

func (r Result[T]) SwitchMap(valueMapper fn.Mapper[T], errorMapper fn.Mapper[error]) Result[T] {
	return r.TrySwitchMap(fn.ToTryableMapper(valueMapper), fn.ToTryableErrorMapper[T](errorMapper))
}

func Combine[T any](results ...Result[T]) Result[[]T] {
	var acc []T
	var errors []error
	lo.ForEach(results,
		func(result Result[T], _ int) {
			if result.IsOk() {
				acc = append(acc, result.GetOrPanic())
			} else {
				errors = append(errors, result.Error())
			}
		})
	if len(errors) == 1 {
		return Error[[]T](errors[0])
	}
	if len(errors) > 1 {
		return Error[[]T](roar.NewAggregatedError(errors...))
	}
	return Ok(acc)
}
