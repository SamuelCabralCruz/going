package fake

import (
	"github.com/SamuelCabralCruz/going/detox/internal/common"
	"github.com/SamuelCabralCruz/going/fn/optional"
	"github.com/samber/lo"
)

type Registration[T any] struct {
	implementation T
	ephemeral      bool
	forCall        optional.Optional[common.Call]
}

func NewRegistration[T any](implementation T, ephemeral bool, forCall optional.Optional[common.Call]) *Registration[T] {
	return &Registration[T]{
		implementation,
		ephemeral,
		forCall,
	}
}

func (i *Registration[T]) IsEphemeral() bool {
	return i.ephemeral
}

func (i *Registration[T]) IsConditional() bool {
	return i.forCall.IsPresent()
}

func (i *Registration[T]) CanHandle(call common.Call) bool {
	if i.IsConditional() {
		return i.forCall.GetOrPanic().EqualTo(call)
	}
	return true
}

func (i *Registration[T]) computePriority() int {
	// Priorities:
	// default implementation (-1/fallback)
	// persistent (0)
	// ephemeral (1)
	// persistentConditional (2)
	// ephemeralConditional (3)
	ephemeralCriterion := lo.If(i.IsEphemeral(), 1).Else(0)
	conditionalCriterion := lo.If(i.IsConditional(), 2).Else(0)
	return conditionalCriterion + ephemeralCriterion
}

func (i *Registration[T]) HasPriorityOver(other *Registration[T]) bool {
	return i.computePriority() > other.computePriority()
}

func (i *Registration[T]) Resolve() T {
	return i.implementation
}
