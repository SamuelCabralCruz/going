package fake

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/samber/lo"
)

type priority int

const (
	persistent priority = iota
	persistentConditional
	ephemeral
	ephemeralConditional
)

type Registration[T any] struct {
	implementation T
	ephemeral      bool
	forCall        optional.Optional[internal.Call]
}

func NewRegistration[T any](implementation T, ephemeral bool, forCall optional.Optional[internal.Call]) *Registration[T] {
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

func (i *Registration[T]) CanHandle(call internal.Call) bool {
	if i.IsConditional() {
		return i.forCall.GetOrPanic().EqualTo(call)
	}
	return true
}

func (i *Registration[T]) computePriority() priority {
	conditionalCriterion := lo.If(i.IsConditional(), 1).Else(0)
	ephemeralCriterion := lo.If(i.IsEphemeral(), 2).Else(0)
	return priority(conditionalCriterion + ephemeralCriterion)
}

func (i *Registration[T]) HasPriorityOver(other *Registration[T]) bool {
	return i.computePriority() > other.computePriority()
}

func (i *Registration[T]) Resolve() T {
	return i.implementation
}
