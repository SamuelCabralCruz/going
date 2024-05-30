package fn

import (
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/SamuelCabralCruz/went/roar"
)

func SafeProducer[T any](producer typing.Producer[T]) (value T, err error) {
	rec := func() {
		if r := recover(); r != nil {
			value = phi.Empty[T]()
			err = roar.AsError(r)
		}
	}
	defer rec()
	if value, err = producer(); err != nil {
		value = phi.Empty[T]()
	}
	return
}

func SafeSupplier[T any](supplier typing.Supplier[T]) (T, error) {
	return SafeProducer(func() (T, error) {
		return supplier(), nil
	})
}

func SafeMapper[T any](mapper typing.Mapper[T], value T) (T, error) {
	return SafeProducer(func() (T, error) {
		return mapper(value), nil
	})
}

func SafeTransformer[T any, U any](transformer typing.Transformer[T, U], value T) (U, error) {
	return SafeProducer(func() (U, error) {
		return transformer(value), nil
	})
}

func SafePredicate[T any](predicate typing.Predicate[T], value T) (bool, error) {
	return SafeProducer(func() (bool, error) {
		return predicate(value), nil
	})
}

func SafeCallable(callable typing.Callable) error {
	_, err := SafeProducer(func() (any, error) {
		callable()
		return nil, nil
	})
	return err
}
