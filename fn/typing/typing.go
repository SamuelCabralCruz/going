package typing

type Callable func()

type Supplier[T any] func() T

type Producer[T any] func() (T, error)

type Consumer[T any] func(T)

type Mapper[T any] func(T) T

type Transformer[T any, U any] func(T) U

type Predicate[T any] func(T) bool

// Asserter
// Prefix: Assert
// Return: assertion (T, error)
type Asserter[T any] func(any) (T, error)

// Validator
// Prefix: Is
// Return: validation (T, ok)
// Should be built on top of asserter
type Validator[T any] func(any) (T, bool)
