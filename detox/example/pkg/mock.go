package pkg

import "github.com/SamuelCabralCruz/went/detox"

// TODO: test without ptr on detox
type AnotherMockClass struct {
	*detox.Detox
}

var _ Another = AnotherMockClass{}

func (m AnotherMockClass) Bye(s string) {
	detox.When(m.Detox, m.Bye).Resolve(s)(s)
}

type SomeMockClass struct {
	*detox.Detox
}

var _ Inter = SomeMockClass{}

func (m SomeMockClass) Hello(s string) (string, error) {
	return detox.When(m.Detox, m.Hello).Resolve(s)(s)

	//// TODO: which dsl is better?
	//mocked := detox.When(m.Detox, m.Hello)
	//mocked.RegisterInvocation(s)
	//return mocked.InvokeFake(s)(s)

	//detox.RegisterInvocation(m.Detox, m.Hello, s)
	//return detox.InvokeFakeImplementation(m.Detox, m.Hello, s)(s)
}

func (m SomeMockClass) Hello2(s string, i int) (string, error, int) {
	return detox.When(m.Detox, m.Hello2).Resolve(s, i)(s, i)
}

func (m SomeMockClass) Prepare() Another {
	return detox.When(m.Detox, m.Prepare).Resolve()()
}
