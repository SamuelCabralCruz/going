package it

import (
	"github.com/SamuelCabralCruz/going/fn"
	"github.com/SamuelCabralCruz/going/fn/typing"
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

func (it *injectable[T]) Clone() InjectionToken[T] {
	return &injectable[T]{
		provider: it.provider,
	}
}
