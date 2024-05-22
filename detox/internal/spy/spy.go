package spy

func NewSpy(mockName string, methodName string) *Spy {
	return &Spy{
		mockName:   mockName,
		methodName: methodName,
	}
}

type Spy struct {
	mockName   string
	methodName string
	calls      [][]any // TODO: add real type? is it necessary?
}

func (s *Spy) RegisterInvocation(args ...any) {
	s.calls = append(s.calls, args)
}

func (s *Spy) Calls() [][]any {
	return s.calls
}

func (s *Spy) CallsCount() int {
	return len(s.Calls())
}

func (s *Spy) NthCall(index int) []any {
	count := s.CallsCount()
	if index >= count {
		panic(newInvalidCallIndexError(s.mockName, s.methodName, index, count))
	}
	return s.Calls()[index]
}
