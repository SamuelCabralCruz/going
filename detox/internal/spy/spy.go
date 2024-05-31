package spy

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/scallop"
)

func NewSpy(info common.MockedMethodInfo) *Spy {
	return &Spy{
		info: info,
	}
}

type Spy struct {
	info  common.MockedMethodInfo
	calls []common.Call
}

func (s *Spy) RegisterCall(call common.Call) {
	s.calls = append(s.calls, call)
}

func (s *Spy) Calls() []common.Call {
	return scallop.Copy(s.calls)
}

func (s *Spy) CallsCount() int {
	return len(s.Calls())
}

func (s *Spy) NthCall(index int) common.Call {
	count := s.CallsCount()
	if index >= count {
		panic(newInvalidCallIndexError(s.info, index, count))
	}
	return s.Calls()[index]
}
