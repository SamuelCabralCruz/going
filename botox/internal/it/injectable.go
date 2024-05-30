package it

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/typing"
)

type injectable[T any] struct {
	provider typing.Producer[T]
}

var _ InjectionToken[struct{}] = &injectable[struct{}]{}

func Register[T any](provider typing.Producer[T]) InjectionToken[T] {
	return &injectable[T]{
		provider: provider,
	}
}

func (it *injectable[T]) Resolve() (T, error) {
	return fn.SafeProducer(it.provider)
}
