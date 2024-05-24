package fake

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
	"sort"
)

func NewFake[T any, U any](info common.MockedMethodInfo, fallback optional.Optional[T]) *Fake[T, U] {
	f := &Fake[T, U]{info: info}
	fallback.IfPresent(f.RegisterDefault)
	return f
}

type Fake[T any, U any] struct {
	info          common.MockedMethodInfo
	fallback      optional.Optional[*Registration[U]]
	registrations []*Registration[U]
}

func (f *Fake[T, U]) RegisterDefault(impl T) {
	method := phi.Value(impl).MethodByName(f.info.Method()).Interface().(U)
	f.fallback = optional.Of(NewRegistration(method, false, optional.Empty[common.Call]()))
}

func (f *Fake[T, U]) Register(impl U) {
	f.register(impl, false, optional.Empty[common.Call]())
}

func (f *Fake[T, U]) RegisterOnce(impl U) {
	f.register(impl, true, optional.Empty[common.Call]())
}

func (f *Fake[T, U]) RegisterConditional(impl U, forCall common.Call) {
	f.register(impl, false, optional.Of(forCall))
}

func (f *Fake[T, U]) RegisterConditionalOnce(impl U, forCall common.Call) {
	f.register(impl, true, optional.Of(forCall))
}

func (f *Fake[T, U]) register(impl U, ephemeral bool, forCall optional.Optional[common.Call]) {
	f.registrations = append(f.registrations, NewRegistration(impl, ephemeral, forCall))
}

func (f *Fake[T, U]) ResolveForCall(call common.Call) U {
	electedCandidate := f.electCandidate(call)
	f.consumeEphemeralCandidate(electedCandidate)
	return electedCandidate.Resolve()
}

func (f *Fake[T, U]) electCandidate(call common.Call) *Registration[U] {
	candidates := f.computeCandidates(call)
	if len(candidates) == 0 {
		return f.resolveFallback()
	}
	candidate := candidates[0]
	return candidate
}

func (f *Fake[T, U]) computeCandidates(call common.Call) []*Registration[U] {
	candidates := lo.Filter(f.registrations, func(item *Registration[U], _ int) bool {
		return item.CanHandle(call)
	})
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].HasPriorityOver(candidates[j])
	})
	return candidates
}

func (f *Fake[T, U]) resolveFallback() *Registration[U] {
	if f.fallback.IsAbsent() {
		panic(newMissingImplementationError(f.info))
	}
	return f.fallback.GetOrPanic()
}

func (f *Fake[T, U]) consumeEphemeralCandidate(candidate *Registration[U]) {
	if candidate.IsEphemeral() {
		f.registrations = lo.Filter(f.registrations, func(item *Registration[U], _ int) bool {
			return item != candidate
		})
	}
}
