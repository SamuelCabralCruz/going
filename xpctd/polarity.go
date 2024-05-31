package xpctd

import "github.com/samber/lo"

type polarity int

const (
	positive polarity = iota
	negative
)

func (p polarity) String() string {
	return lo.Ternary(p == negative, "not to", "to")
}
