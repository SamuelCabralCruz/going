package faked

func NewUnconditionalFaked[T any](impl T) *UnconditionalFaked[T] {
	return &UnconditionalFaked[T]{impl}
}

type UnconditionalFaked[T any] struct {
	impl T
}

var _ Faked[any] = &UnconditionalFaked[any]{}

func (f *UnconditionalFaked[T]) CanHandle(_ ...any) bool {
	return true
}

func (f *UnconditionalFaked[T]) Invoke() T {
	return f.impl
}
