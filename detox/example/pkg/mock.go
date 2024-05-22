package pkg

import "github.com/SamuelCabralCruz/went/detox"

// TODO: test without ptr on detox
type AnotherMockClass struct {
	*detox.Detox
}

var _ Another = AnotherMockClass{}

func (m AnotherMockClass) Bye(s string) {
	// TODO: should be wrapper inside
	detox.RegisterInvocation(m.Detox, m.Bye, s)
	detox.InvokeFakeImplementation(m.Detox, m.Bye)(s)
}

type SomeMockClass struct {
	*detox.Detox
}

var _ Inter = SomeMockClass{}

func (m SomeMockClass) Hello(s string) (string, error) {
	detox.RegisterInvocation(m.Detox, m.Hello, s)
	return detox.InvokeFakeImplementation(m.Detox, m.Hello)(s)
}

func (m SomeMockClass) Prepare() Another {
	detox.RegisterInvocation(m.Detox, m.Prepare)
	return detox.InvokeFakeImplementation(m.Detox, m.Prepare)()
}
