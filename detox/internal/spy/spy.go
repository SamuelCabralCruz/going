package spy

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
	"github.com/SamuelCabralCruz/went/fn"
)

func NewSpy(mockName string, methodName string) *Spy {
	return &Spy{
		mockName:   mockName,
		methodName: methodName,
	}
}

type Spy struct {
	mockName   string
	methodName string
	calls      []internal.Call
}

func (s *Spy) RegisterCall(call internal.Call) {
	s.calls = append(s.calls, call)
}

func (s *Spy) Calls() []internal.Call {
	return fn.Copy(s.calls)
}

func (s *Spy) CallsCount() int {
	return len(s.Calls())
}

func (s *Spy) NthCall(index int) internal.Call {
	count := s.CallsCount()
	if index >= count {
		panic(newInvalidCallIndexError(s.mockName, s.methodName, index, count))
	}
	return s.Calls()[index]
}
