package detox

func When[T any](mock *Detox, orig T) Mocked[T] {
	return Mocked[T]{
		mock: mock,
		orig: orig,
	}
}

type Mocked[T any] struct {
	mock *Detox
	orig T
}

func (m Mocked[T]) Call(impl T) {
	FakeImplementation(m.mock, m.orig, impl)
}

func (m Mocked[T]) CallOnce(impl T) {
	FakeImplementationOnce(m.mock, m.orig, impl)
}

// TODO: make private
func (m Mocked[T]) RegisterInvocation(args ...any) {
	RegisterInvocation(m.mock, m.orig, args...)
}

// TODO: make private
func (m Mocked[T]) InvokeFake(args ...any) T {
	return InvokeFakeImplementation(m.mock, m.orig, args...)
}

func (m Mocked[T]) Resolve(args ...any) T {
	m.RegisterInvocation(args...)
	return m.InvokeFake(args...)
}

func (m Mocked[T]) Calls() [][]any {
	return Calls(m.mock, m.orig)
}

func (m Mocked[T]) CallsCount() int {
	return CallsCount(m.mock, m.orig)
}

func (m Mocked[T]) NthCall(index int) []any {
	return NthCall(m.mock, m.orig, index)
}

func (m Mocked[T]) Reset() {
	Reset(m.mock, m.orig)
}

type MockedConditionally[T any] struct {
	Mocked[T]
	withArgs []any
}

func (m Mocked[T]) WithArgs(args ...any) MockedConditionally[T] {
	return MockedConditionally[T]{
		Mocked:   m,
		withArgs: args,
	}
}

func (m MockedConditionally[T]) Call(impl T) {
	FakeConditionalImplementation(m.mock, m.orig, impl, m.withArgs)
}

func (m MockedConditionally[T]) CallOnce(impl T) {
	FakeConditionalImplementationOnce(m.mock, m.orig, impl, m.withArgs)
}
