package fn

// Producer There might be error at production time
type Producer[T any] func() (T, error)

// Supplier Products being shipped have already passed quality insurance
type Supplier[T any] func() T

type Consumer[T any] func(T)
