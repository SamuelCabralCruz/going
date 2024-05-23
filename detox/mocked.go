package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
)

func When[T any, U any](mock *Detox[T], method U) Mocked[U] {
	return &mocked[T, U]{
		info:   common.NewMockedMethodInfo(mock.info, method),
		mock:   mock,
		method: method,
	}
}

type Mocked[T any] interface {
	ResolveForArgs(...any) T
	Call(T)
	CallOnce(T)
	Reset()
	WithArgs(...any) MockedConditionally[T]
	Assertable
}

type mocked[T any, U any] struct {
	info   common.MockedMethodInfo
	mock   *Detox[T]
	method U
}

func (m *mocked[T, U]) Describe() string {
	return m.info.Describe()
}

func (m *mocked[T, U]) ResolveForArgs(args ...any) U {
	call := common.NewCall(args...)
	resolveSpy(m.mock, m.method).RegisterCall(call)
	return resolveFake(m.mock, m.method).ResolveForCall(call)
}

func (m *mocked[T, U]) Call(implementation U) {
	resolveFake(m.mock, m.method).RegisterImplementation(implementation)
}

func (m *mocked[T, U]) CallOnce(implementation U) {
	resolveFake(m.mock, m.method).RegisterImplementationOnce(implementation)
}

func (m *mocked[T, U]) Reset() {
	id := getId(m.method)
	delete(m.mock.fakes, id)
	delete(m.mock.spies, id)
}

func (m *mocked[T, U]) Calls() []common.Call {
	return resolveSpy(m.mock, m.method).Calls()
}

func (m *mocked[T, U]) CallsCount() int {
	return resolveSpy(m.mock, m.method).CallsCount()
}

func (m *mocked[T, U]) NthCall(index int) common.Call {
	return resolveSpy(m.mock, m.method).NthCall(index)
}
