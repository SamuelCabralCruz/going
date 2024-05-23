package fake

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/samber/lo"
	"sort"
)

func NewFake[T any](info common.MockedMethodInfo, fallback optional.Optional[T]) *Fake[T] {
	return &Fake[T]{
		info: info,
		fallback: optional.Transform(fallback, func(f T) *Registration[T] {
			return NewRegistration(f, false, optional.Empty[common.Call]())
		}),
	}
}

type Fake[T any] struct {
	info          common.MockedMethodInfo
	fallback      optional.Optional[*Registration[T]]
	registrations []*Registration[T]
}

func (f *Fake[T]) RegisterImplementation(impl T) {
	f.register(impl, false, optional.Empty[common.Call]())
}

func (f *Fake[T]) RegisterImplementationOnce(impl T) {
	f.register(impl, true, optional.Empty[common.Call]())
}

func (f *Fake[T]) RegisterConditionalImplementation(impl T, forCall common.Call) {
	f.register(impl, false, optional.Of(forCall))
}

func (f *Fake[T]) RegisterConditionalImplementationOnce(impl T, forCall common.Call) {
	f.register(impl, true, optional.Of(forCall))
}

func (f *Fake[T]) register(implementation T, ephemeral bool, forCall optional.Optional[common.Call]) {
	f.registrations = append(f.registrations, NewRegistration(implementation, ephemeral, forCall))
}

func (f *Fake[T]) ResolveForCall(call common.Call) T {
	electedCandidate := f.electCandidate(call)
	f.consumeEphemeralCandidate(electedCandidate)
	return electedCandidate.Resolve()
}

func (f *Fake[T]) electCandidate(call common.Call) *Registration[T] {
	candidates := f.computeCandidates(call)
	if len(candidates) == 0 {
		return f.resolveFallback()
	}
	candidate := candidates[0]
	return candidate
}

func (f *Fake[T]) computeCandidates(call common.Call) []*Registration[T] {
	candidates := lo.Filter(f.registrations, func(item *Registration[T], _ int) bool {
		return item.CanHandle(call)
	})
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].HasPriorityOver(candidates[j])
	})
	return candidates
}

func (f *Fake[T]) resolveFallback() *Registration[T] {
	if f.fallback.IsAbsent() {
		panic(newMissingImplementationError(f.info))
	}
	return f.fallback.GetOrPanic()
}

func (f *Fake[T]) consumeEphemeralCandidate(candidate *Registration[T]) {
	if candidate.IsEphemeral() {
		f.registrations = lo.Filter(f.registrations, func(item *Registration[T], _ int) bool {
			return item != candidate
		})
	}
}
