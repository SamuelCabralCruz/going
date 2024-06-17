//go:build test

package fixture

import (
	"github.com/SamuelCabralCruz/going/detox"
)

type Interface3 interface {
	Method(string)
	AnotherMethod(int)
}

type Implementation3 struct {
	Called bool
}

var _ Interface3 = &Implementation3{}

func (i *Implementation3) Method(_ string) {
	i.Called = true
}

func (i *Implementation3) AnotherMethod(_ int) {}

func NewInterface3Mock() Interface3Mock {
	return Interface3Mock{detox.New[Interface3]()}
}

type Interface3Mock struct {
	*detox.Detox[Interface3]
}

var _ Interface3 = Interface3Mock{}

func (i Interface3Mock) Method(s string) {
	detox.When(i.Detox, i.Method).ResolveForArgs(s)(s)
}

func (i Interface3Mock) AnotherMethod(s int) {
	detox.When(i.Detox, i.AnotherMethod).ResolveForArgs(s)(s)
}
