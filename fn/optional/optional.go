package optional

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple"
	"github.com/SamuelCabralCruz/went/phi"
)

func Empty[T any]() Optional[T] {
	return Optional[T]{
		isPresent: false,
	}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{
		value:     value,
		isPresent: true,
	}
}

func OfNullable[T any](value T) Optional[T] {
	if phi.IsZero(value) {
		return Empty[T]()
	}
	return Of(value)
}

func FromTuple[T any](value T, _ error) Optional[T] {
	return OfNullable(value)
}

func FromSupplier[T any](supply fn.Supplier[T]) Optional[T] {
	return FromTuple(fn.Safe(supply))
}

func FromProducer[T any](produce fn.Producer[T]) Optional[T] {
	return FromTuple(fn.Try(produce))
}

type Optional[T any] struct {
	value     T
	isPresent bool
}

func (o Optional[T]) IsPresent() bool {
	return o.isPresent
}

func (o Optional[T]) IsAbsent() bool {
	return !o.IsPresent()
}

func (o Optional[T]) Get() (T, error) {
	if o.IsAbsent() {
		return o.value, newNoSuchElementError()
	}
	return o.value, nil
}

func (o Optional[T]) GetOrPanic() T {
	return tuple.GetOrPanic(o.Get())
}

func (o Optional[T]) GetOrPanicWith(err error) T {
	if o.IsAbsent() {
		panic(err)
	}
	return o.value
}

func (o Optional[T]) OrEmpty() T {
	return o.value
}

func (o Optional[T]) OrElseTry(produce fn.Producer[T]) Optional[T] {
	if o.IsAbsent() {
		return FromProducer(produce)
	}
	return o
}

func (o Optional[T]) OrElseGet(supply fn.Supplier[T]) T {
	return o.OrElseTry(fn.ToProducer(supply)).OrEmpty()
}

func (o Optional[T]) OrElse(value T) T {
	return o.OrElseGet(fn.ToSupplier(value))
}

func (o Optional[T]) TryFlatMap(mapper fn.TryableMapper[Optional[T]]) Optional[T] {
	if o.IsPresent() {
		maybe, err := mapper(o)
		if err != nil {
			return Empty[T]()
		}
		return maybe
	}
	return o
}

func (o Optional[T]) TryFlatMapEmpty(produce fn.Producer[Optional[T]]) Optional[T] {
	if o.IsAbsent() {
		maybe, err := produce()
		if err != nil {
			return Empty[T]()
		}
		return maybe
	}
	return o
}

func (o Optional[T]) FlatMap(mapper fn.Mapper[Optional[T]]) Optional[T] {
	return o.TryFlatMap(fn.ToTryableMapper(mapper))
}

func (o Optional[T]) FlatMapEmpty(supply fn.Supplier[Optional[T]]) Optional[T] {
	return o.TryFlatMapEmpty(fn.ToProducer(supply))
}

func (o Optional[T]) TryMap(mapper fn.TryableMapper[T]) Optional[T] {
	if o.IsPresent() {
		return FromTuple[T](mapper(o.value))
	}
	return o
}

func (o Optional[T]) TryMapEmpty(produce fn.Producer[T]) Optional[T] {
	if o.IsAbsent() {
		return FromTuple[T](produce())
	}
	return o
}

func (o Optional[T]) Map(mapper fn.Mapper[T]) Optional[T] {
	return o.TryMap(fn.ToTryableMapper(mapper))
}

func (o Optional[T]) MapEmpty(supply fn.Supplier[T]) Optional[T] {
	return o.TryMapEmpty(fn.ToProducer(supply))
}

func (o Optional[T]) IfPresent(consume fn.Consumer[T]) {
	if o.IsPresent() {
		consume(o.value)
	}
}

func (o Optional[T]) Filter(predict fn.Predicate[T]) Optional[T] {
	if o.IsPresent() && predict(o.value) {
		return o
	}
	return Empty[T]()
}
