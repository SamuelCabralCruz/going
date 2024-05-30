package optional

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/fn/tuple/validation"
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/phi"
)

func Empty[T any]() Optional[T] {
	return Optional[T]{
		isPresent: false,
	}
}

func ofError[T any](_ error) Optional[T] {
	return Empty[T]()
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

func FromAssertion[T any](value T, err error) Optional[T] {
	if err != nil {
		return Empty[T]()
	}
	return OfNullable(value)
}

func FromValidation[T any](value T, ok bool) Optional[T] {
	return FromAssertion(validation.ToAssertion(value, ok))
}

func FromSupplier[T any](supplier typing.Supplier[T]) Optional[T] {
	return FromAssertion(fn.SafeSupplier(supplier))
}

func FromProducer[T any](producer typing.Producer[T]) Optional[T] {
	return FromAssertion(fn.SafeProducer(producer))
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
		return phi.Empty[T](), newNoSuchElementError()
	}
	return o.value, nil
}

func (o Optional[T]) GetOrPanic() T {
	return assertion.GetOrPanic(o.Get())
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

func (o Optional[T]) OrElseGet(supplier typing.Supplier[T]) T {
	if o.IsAbsent() {
		return FromSupplier(supplier).OrEmpty()
	}
	return o.value
}

func (o Optional[T]) OrElse(value T) T {
	if o.IsAbsent() {
		return value
	}
	return o.value
}

func (o Optional[T]) FlatMap(mapper typing.Mapper[Optional[T]]) Optional[T] {
	if o.IsPresent() {
		return assertion.Switch[Optional[T], Optional[T]](fn.SafeMapper(mapper, o))(fn.Identity[Optional[T]], ofError[T])
	}
	return o
}

func (o Optional[T]) FlatMapEmpty(supplier typing.Supplier[Optional[T]]) Optional[T] {
	if o.IsAbsent() {
		return assertion.Switch[Optional[T], Optional[T]](fn.SafeSupplier(supplier))(fn.Identity[Optional[T]], ofError[T])
	}
	return o
}

func (o Optional[T]) FlatSwitchMap(onPresent typing.Mapper[Optional[T]], onAbsent typing.Supplier[Optional[T]]) Optional[T] {
	if o.IsPresent() {
		return o.FlatMap(onPresent)
	}
	return o.FlatMapEmpty(onAbsent)
}

func (o Optional[T]) Map(mapper typing.Mapper[T]) Optional[T] {
	if o.IsPresent() {
		return assertion.Switch[T, Optional[T]](fn.SafeMapper(mapper, o.value))(OfNullable[T], ofError[T])
	}
	return o
}

func (o Optional[T]) MapEmpty(supplier typing.Supplier[T]) Optional[T] {
	if o.IsAbsent() {
		return assertion.Switch[T, Optional[T]](fn.SafeSupplier(supplier))(OfNullable[T], ofError[T])
	}
	return o
}

func (o Optional[T]) SwitchMap(onPresent typing.Mapper[T], onAbsent typing.Supplier[T]) Optional[T] {
	if o.IsPresent() {
		return o.Map(onPresent)
	}
	return o.MapEmpty(onAbsent)
}

func (o Optional[T]) IfPresent(consumer typing.Consumer[T]) {
	if o.IsPresent() {
		consumer(o.value)
	}
}

func (o Optional[T]) Filter(predicate typing.Predicate[T]) Optional[T] {
	if o.IsPresent() {
		filterWithPredicate := func(predicated bool) Optional[T] {
			if predicated {
				return o
			}
			return Empty[T]()
		}
		assertion.Switch[bool, Optional[T]](fn.SafePredicate(predicate, o.value))(filterWithPredicate, ofError[T])
	}
	return o
}
