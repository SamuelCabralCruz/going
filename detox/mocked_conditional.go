package detox

import "github.com/SamuelCabralCruz/went/detox/internal"

func (m *mocked[T]) WithArgs(args ...any) MockedConditionally[T] {
	return &mockedConditionally[T]{
		m,
		internal.NewCall(args...),
	}
}

type MockedConditionally[T any] interface {
	Call(T)
	CallOnce(T)
}

type mockedConditionally[T any] struct {
	*mocked[T]
	forCall internal.Call
}

func (m *mockedConditionally[T]) Call(impl T) {
	resolveFake(m.mock, m.orig).RegisterConditionalImplementation(impl, m.forCall)
}

func (m *mockedConditionally[T]) CallOnce(impl T) {
	resolveFake(m.mock, m.orig).RegisterConditionalImplementationOnce(impl, m.forCall)
}
