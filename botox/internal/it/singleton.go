package it

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/result"
)

type singleton[T any] struct {
	token     InjectionToken[T]
	provided  bool
	reference result.Result[T]
}

var _ InjectionToken[struct{}] = &singleton[struct{}]{}

func RegisterSingleton[T any](provider fn.Producer[T]) InjectionToken[T] {
	return &singleton[T]{
		token: &injectable[T]{
			provider: provider,
		},
	}
}

func (s *singleton[T]) Resolve() (T, error) {
	if s.provided {
		return s.reference.Get()
	}
	s.provided = true
	s.reference = result.FromTuple(s.token.Resolve())
	return s.reference.Get()
}

func (s *singleton[T]) MustResolve() T {
	return s.reference.GetOrPanic()
}
