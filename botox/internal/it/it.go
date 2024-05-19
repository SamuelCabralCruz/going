package it

type InjectionToken[T any] interface {
	Resolve() (T, error)
	MustResolve() T
}
