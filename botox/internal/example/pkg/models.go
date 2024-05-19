package pkg

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/samber/lo"
)

type Test1 string
type Test2 string

type SomeInterface interface {
	Coucou() string
}

type SomeImpl1 struct {
}

func (i SomeImpl1) Coucou() string {
	return "impl1"
}

type SomeImpl2 struct {
}

func (i SomeImpl2) Coucou() string {
	return "impl2"
}

type SomeImpl3 struct {
}

func (i SomeImpl3) Coucou() string {
	return "impl3"
}

type Nested struct {
	inter []SomeInterface
	t1    Test1
	t2    Test2
}

func NewNested() Nested {
	return Nested{
		inter: botox.MustResolveAll[SomeInterface](),
		t1:    botox.MustResolve[Test1](),
		t2:    botox.MustResolve[Test2](),
	}
}

func (n Nested) ToString() string {
	return fmt.Sprintf("Nested t1: %s t2: %s inter: %s", n.t1, n.t2, lo.Map(n.inter, func(item SomeInterface, _ int) string {
		return item.Coucou()
	}))
}
