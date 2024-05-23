package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
)

func When[T any](mock *Detox, orig T) Mocked[T] {
	return &mocked[T]{
		mock: mock,
		orig: orig,
	}
}

type Mocked[T any] interface {
	Name() string

	ResolveForArgs(...any) T
	Call(T)
	CallOnce(T)
	Reset()

	WithArgs(...any) MockedConditionally[T]
	Assert() Asserter
}

type mocked[T any] struct {
	mock *Detox
	orig T
}

func (m *mocked[T]) Name() string {
	return name(m.mock, m.orig)
}

func (m *mocked[T]) ResolveForArgs(args ...any) T {
	call := internal.NewCall(args...)
	resolveSpy(m.mock, m.orig).RegisterCall(call)
	return resolveFake(m.mock, m.orig).ResolveForCall(call)
}

func (m *mocked[T]) Call(implementation T) {
	resolveFake(m.mock, m.orig).RegisterImplementation(implementation)
}

func (m *mocked[T]) CallOnce(implementation T) {
	resolveFake(m.mock, m.orig).RegisterImplementationOnce(implementation)
}

func (m *mocked[T]) Reset() {
	id := getId(m.orig)
	delete(m.mock.fakes, id)
	delete(m.mock.spies, id)
}

func (m *mocked[T]) Calls() []internal.Call {
	return resolveSpy(m.mock, m.orig).Calls()
}

func (m *mocked[T]) CallsCount() int {
	return resolveSpy(m.mock, m.orig).CallsCount()
}

func (m *mocked[T]) NthCall(index int) internal.Call {
	return resolveSpy(m.mock, m.orig).NthCall(index)
}
