package fn

// Supplier Products being shipped have already passed quality insurance
type Supplier[T any] func() T

// Producer There might be error at production time
type Producer[T any] func() (T, error)

type Consumer[T any] func(T)

type Predicate[T any] func(T) bool

type Mapper[T any] func(T) T

type TryableMapper[T any] func(T) (T, error)

type TryableErrorMapper[T any] func(error) (T, error)

// TODO: remove?
//type TupleMapper[T any] func(T, error) (T, error)

type Transformer[T any, U any] func(T) U

type TryableTransformer[T any, U any] func(T) (U, error)
