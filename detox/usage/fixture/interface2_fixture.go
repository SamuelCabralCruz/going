//go:build test

package fixture

import (
	"github.com/SamuelCabralCruz/going/detox"
)

type Interface2 interface {
	ReturnAnotherInterface() Interface1
	AnotherMethod(string) string
}

type Implementation2 struct{}

var _ Interface2 = Implementation2{}

func (i Implementation2) ReturnAnotherInterface() Interface1 {
	return Implementation1{}
}

func (i Implementation2) AnotherMethod(_ string) string {
	return "ret1"
}

func NewInterface2Mock() Interface2Mock {
	return Interface2Mock{detox.New[Interface2]()}
}

type Interface2Mock struct {
	*detox.Detox[Interface2]
}

var _ Interface2 = Interface2Mock{}

func (i Interface2Mock) ReturnAnotherInterface() Interface1 {
	return detox.When(i.Detox, i.ReturnAnotherInterface).ResolveForArgs()()
}

func (i Interface2Mock) AnotherMethod(s string) string {
	return detox.When(i.Detox, i.AnotherMethod).ResolveForArgs(s)(s)
}
