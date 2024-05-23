package fake

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/samber/lo"
	"sort"
)

func NewFake[T any](mockName string, methodName string) *Fake[T] {
	return &Fake[T]{
		mockName:   mockName,
		methodName: methodName,
	}
}

type Fake[T any] struct {
	mockName      string
	methodName    string
	registrations []*Registration[T]
}

func (f *Fake[T]) RegisterImplementation(impl T) {
	f.register(impl, false, optional.Empty[internal.Call]())
}

func (f *Fake[T]) RegisterImplementationOnce(impl T) {
	f.register(impl, true, optional.Empty[internal.Call]())
}

func (f *Fake[T]) RegisterConditionalImplementation(impl T, forCall internal.Call) {
	f.register(impl, false, optional.Of(forCall))
}

func (f *Fake[T]) RegisterConditionalImplementationOnce(impl T, forCall internal.Call) {
	f.register(impl, true, optional.Of(forCall))
}

func (f *Fake[T]) register(implementation T, ephemeral bool, forCall optional.Optional[internal.Call]) {
	f.registrations = append(f.registrations, NewRegistration(implementation, ephemeral, forCall))
}

func (f *Fake[T]) ResolveForCall(call internal.Call) T {
	candidates := lo.Filter(f.registrations, func(item *Registration[T], _ int) bool {
		return item.CanHandle(call)
	})

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].HasPriorityOver(candidates[j])
	})

	if len(candidates) == 0 {
		panic(newMissingFakeImplementationError(f.mockName, f.methodName))
	}
	candidate := candidates[0]

	if candidate.IsEphemeral() {
		f.registrations = lo.Filter(f.registrations, func(item *Registration[T], _ int) bool {
			return item != candidate
		})
	}

	return candidate.Resolve()
}
