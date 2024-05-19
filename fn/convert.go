package fn

import "github.com/SamuelCabralCruz/went/phi"

func ToSupplier[T any](value T) Supplier[T] {
	return func() T {
		return value
	}
}

func ToMapper[T any](supply Supplier[T]) Mapper[T] {
	return func(_ T) T {
		return supply()
	}
}

func ToProducer[T any](supply Supplier[T]) Producer[T] {
	return func() (T, error) {
		return Safe(supply)
	}
}

func ToTryableMapper[T any](mapper Mapper[T]) TryableMapper[T] {
	return func(value T) (T, error) {
		return Safe(func() T {
			return mapper(value)
		})
	}
}

func ToTryableErrorMapper[T any](mapper Mapper[error]) TryableErrorMapper[T] {
	return func(err error) (T, error) {
		return Try(func() (T, error) {
			return phi.Empty[T](), mapper(err)
		})
	}
}

func ToTryableTransformer[T any, U any](transformer Transformer[T, U]) TryableTransformer[T, U] {
	return func(value T) (U, error) {
		return Safe(func() U {
			return transformer(value)
		})
	}
}

// TODO: remove?
//func ToTupleMapper[T any](produce Producer[T]) TupleMapper[T] {
//	return func(_ T, err1 error) (T, error) {
//		v2, err2 := Try(produce)
//		return v2, roar.Combine(err1, err2)
//	}
//}
