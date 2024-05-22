package faked

type Faked[T any] interface {
	CanHandle(...any) bool
	Invoke() T
}
