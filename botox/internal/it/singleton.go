package it

import (
	"github.com/samber/mo"
)

type singleton[T any] struct {
	token     InjectionToken[T]
	provided  bool
	reference mo.Result[T]
}

var _ InjectionToken[struct{}] = &singleton[struct{}]{}

func RegisterSingleton[T any](provider Provider[T]) InjectionToken[T] {
	return &singleton[T]{
		token: &injectable[T]{
			provider: provider,
		},
	}
}

func (s *singleton[T]) Resolve() mo.Result[T] {
	if s.provided {
		return s.reference
	}
	s.provided = true
	s.reference = s.token.Resolve()
	return s.reference
}

func (s *singleton[T]) MustResolve() T {
	value, err := s.Resolve().Get()
	if err != nil {
		panic(err)
	}
	return value
}
