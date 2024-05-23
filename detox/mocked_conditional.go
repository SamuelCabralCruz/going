package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
)

func (m *mocked[T, U]) WithArgs(args ...any) MockedConditionally[U] {
	return &mockedConditionally[T, U]{
		m,
		common.NewCall(args...),
	}
}

type MockedConditionally[T any] interface {
	Call(T)
	CallOnce(T)
}

type mockedConditionally[T any, U any] struct {
	*mocked[T, U]
	forCall common.Call
}

func (m *mockedConditionally[T, U]) Call(impl U) {
	resolveFake(m.mock, m.method).RegisterConditionalImplementation(impl, m.forCall)
}

func (m *mockedConditionally[T, U]) CallOnce(impl U) {
	resolveFake(m.mock, m.method).RegisterConditionalImplementationOnce(impl, m.forCall)
}
