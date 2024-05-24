//go:build test

package usage

import (
	"github.com/SamuelCabralCruz/went/detox"
)

type Interface3 interface {
	Method()
}

type Implementation3 struct {
	Called bool
}

var _ Interface3 = &Implementation3{}

func (i *Implementation3) Method() {
	i.Called = true
}

func NewInterface3Mock() Interface3Mock {
	return Interface3Mock{detox.New[Interface3]()}
}

type Interface3Mock struct {
	*detox.Detox[Interface3]
}

var _ Interface3 = Interface3Mock{}

func (i Interface3Mock) Method() {
	detox.When(i.Detox, i.Method).ResolveForArgs()()
}