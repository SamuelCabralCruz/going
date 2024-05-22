package internal

func NewSpy() *Spy {
	return &Spy{}
}

// TODO: add real type
type Spy struct {
	calls [][]any
}

func (s *Spy) RegisterInvocation(args ...any) {
	s.calls = append(s.calls, args)
}

func (s *Spy) Count() int {
	return len(s.calls)
}

func (s *Spy) Calls() [][]any {
	return s.calls
}
