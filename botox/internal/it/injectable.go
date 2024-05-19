package it

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/result"
)

type injectable[T any] struct {
	provider fn.Producer[T]
}

var _ InjectionToken[struct{}] = &injectable[struct{}]{}

func Register[T any](provider fn.Producer[T]) InjectionToken[T] {
	return &injectable[T]{
		provider: provider,
	}
}

func (it *injectable[T]) Resolve() (T, error) {
	return it.provider()
}

func (it *injectable[T]) MustResolve() T {
	return result.FromTuple(it.Resolve()).GetOrPanic()
}
