package fn

import (
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/fn/tuple/validation"
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/phi"
)

func ValueToSupplier[T any](value T) typing.Supplier[T] {
	return func() T {
		return value
	}
}

func SupplierToProducer[T any](supplier typing.Supplier[T]) typing.Producer[T] {
	return func() (T, error) {
		return supplier(), nil
	}
}

func ProducerToSupplier[T any](producer typing.Producer[T]) typing.Supplier[T] {
	return func() T {
		return assertion.GetOrPanic(producer())
	}
}

func SupplierToMapper[T any](supplier typing.Supplier[T]) typing.Mapper[T] {
	return func(_ T) T {
		return supplier()
	}
}

func MapperToSupplier[T any](mapper typing.Mapper[T], value T) func() T {
	return func() T {
		return mapper(value)
	}
}

func MapperToTransformer[T any](mapper typing.Mapper[T]) typing.Transformer[T, T] {
	return func(value T) T {
		return mapper(value)
	}
}

func AsserterToValidator[T any](asserter typing.Asserter[T]) typing.Validator[T] {
	return func(value any) (T, bool) {
		asserted, err := asserter(value)
		return asserted, err == nil
	}
}

func ValidatorToPredicate[T any](validator typing.Validator[T]) typing.Predicate[any] {
	return func(value any) bool {
		_, ok := validator(value)
		return ok
	}
}

func PredicateToValidator[T any](predicate typing.Predicate[T]) typing.Validator[T] {
	return func(value any) (T, bool) {
		v, ok := value.(T)
		if ok && predicate(v) {
			return v, true
		}
		return phi.Empty[T](), false
	}
}

func ValidatorToAsserter[T any](validator typing.Validator[T]) typing.Asserter[T] {
	return func(value any) (T, error) {
		return validation.ToAssertion(validator(value))
	}
}
