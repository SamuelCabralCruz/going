package pkg

import "github.com/SamuelCabralCruz/went/detox"

func NewAnotherMockClass() AnotherMockClass {
	return AnotherMockClass{detox.New[Another]()}
}

type AnotherMockClass struct {
	*detox.Detox[Another]
}

var _ Another = AnotherMockClass{}

func (m AnotherMockClass) Bye(s string) {
	detox.When(m.Detox, m.Bye).ResolveForArgs(s)(s)
}

func NewSomeMockClass() SomeMockClass {
	return SomeMockClass{detox.New[Inter]()}
}

type SomeMockClass struct {
	*detox.Detox[Inter]
}

var _ Inter = SomeMockClass{}

func (m SomeMockClass) Hello(s string) (string, error) {
	return detox.When(m.Detox, m.Hello).ResolveForArgs(s)(s)
}

func (m SomeMockClass) Hello2(s string, i int) (string, error, int) {
	return detox.When(m.Detox, m.Hello2).ResolveForArgs(s, i)(s, i)
}

func (m SomeMockClass) Prepare() Another {
	return detox.When(m.Detox, m.Prepare).ResolveForArgs()()
}
