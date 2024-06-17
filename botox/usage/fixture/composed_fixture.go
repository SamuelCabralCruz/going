//go:build test

package fixture

import (
	"fmt"
	"github.com/SamuelCabralCruz/going/botox"
	"strings"
)

func NewComposed() Composed {
	return Composed{
		multiple:  botox.MustResolveAll[SomeInterface](),
		stateless: botox.MustResolve[Stateless](),
		Stateful:  botox.MustResolve[*Stateful](),
	}
}

type Composed struct {
	multiple  []SomeInterface
	stateless Stateless
	Stateful  *Stateful
}

func (c Composed) Method() string {
	var multipleParts []string
	for _, d := range c.multiple {
		multipleParts = append(multipleParts, d.Method())
	}
	return fmt.Sprintf("multiple: %s stateless: %s stateful %s",
		fmt.Sprintf("[%s]", strings.Join(multipleParts, ", ")),
		c.stateless.Method(),
		c.Stateful.Method())
}
