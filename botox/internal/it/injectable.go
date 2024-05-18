package it

import "github.com/samber/mo"

type injectable[T any] struct {
	provider Provider[T]
}

var _ InjectionToken[struct{}] = &injectable[struct{}]{}

func Register[T any](provider Provider[T]) InjectionToken[T] {
	return &injectable[T]{
		provider: provider,
	}
}

func (it *injectable[T]) Resolve() mo.Result[T] {
	return it.provider()
}

func (it *injectable[T]) MustResolve() T {
	value, err := it.Resolve().Get()
	if err != nil {
		panic(err)
	}
	return value
}
