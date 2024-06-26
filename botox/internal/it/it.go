package it

type InjectionToken[T any] interface {
	Resolve() (T, error)
	Clone() InjectionToken[T]
}
