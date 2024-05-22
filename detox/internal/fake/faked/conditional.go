package faked

import "reflect"

func NewConditionalFaked[T any](impl T, withArgs []any) *ConditionalFaked[T] {
	return &ConditionalFaked[T]{impl, withArgs}
}

type ConditionalFaked[T any] struct {
	impl     T
	withArgs []any
}

var _ Faked[any] = &ConditionalFaked[any]{}

func (f *ConditionalFaked[T]) CanHandle(args ...any) bool {
	return reflect.DeepEqual(args, f.withArgs)
}

func (f *ConditionalFaked[T]) Invoke() T {
	return f.impl
}
