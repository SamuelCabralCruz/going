package fake

import (
	"github.com/SamuelCabralCruz/going/detox/internal/common"
	"github.com/SamuelCabralCruz/going/fn/optional"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/samber/lo"
	"sort"
)

func NewFake[T any, U any](info common.MockedMethodInfo, fallback optional.Optional[T]) *Fake[T, U] {
	f := &Fake[T, U]{info: info}
	fallback.IfPresent(f.RegisterDefault)
	return f
}

type Fake[T any, U any] struct {
	info                   common.MockedMethodInfo
	fallback               optional.Optional[*Registration[U]]
	persistent             optional.Optional[*Registration[U]]
	persistentConditionals []*Registration[U]
	ephemerals             []*Registration[U]
}

func (f *Fake[T, U]) RegisterDefault(impl T) {
	method := phi.Value(impl).MethodByName(f.info.Method()).Interface().(U)
	f.fallback = optional.Of(NewRegistration(method, false, optional.Empty[common.Call]()))
}

func (f *Fake[T, U]) Register(impl U) {
	f.persistent = optional.Of(NewRegistration(impl, false, optional.Empty[common.Call]()))
}

func (f *Fake[T, U]) RegisterOnce(impl U) {
	f.ephemerals = append(f.ephemerals, NewRegistration(impl, true, optional.Empty[common.Call]()))
}

func (f *Fake[T, U]) RegisterConditional(impl U, forCall common.Call) {
	f.persistentConditionals = append(
		lo.Filter(f.persistentConditionals, func(item *Registration[U], _ int) bool {
			return !item.CanHandle(forCall)
		}),
		NewRegistration(impl, false, optional.Of(forCall)))
}

func (f *Fake[T, U]) RegisterConditionalOnce(impl U, forCall common.Call) {
	f.ephemerals = append(f.ephemerals, NewRegistration(impl, true, optional.Of(forCall)))
}

func (f *Fake[T, U]) ResolveForCall(call common.Call) U {
	electedCandidate := f.electCandidate(call)
	f.consumeEphemeralCandidate(electedCandidate)
	return electedCandidate.Resolve()
}

func (f *Fake[T, U]) electCandidate(call common.Call) *Registration[U] {
	candidates := f.computeCandidates(call)
	if len(candidates) == 0 {
		return f.resolveFallback(call)
	}
	candidate := candidates[0]
	return candidate
}

func (f *Fake[T, U]) computeCandidates(call common.Call) []*Registration[U] {
	candidates := lo.Filter(f.getAllRegistrations(), func(item *Registration[U], _ int) bool {
		return item.CanHandle(call)
	})
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].HasPriorityOver(candidates[j])
	})
	return candidates
}

func (f *Fake[T, U]) getAllRegistrations() []*Registration[U] {
	return lo.Flatten([][]*Registration[U]{
		optional.Transform(
			f.persistent,
			func(r *Registration[U]) []*Registration[U] {
				return []*Registration[U]{r}
			}).OrElse([]*Registration[U]{}),
		f.ephemerals,
		f.persistentConditionals,
	})
}

func (f *Fake[T, U]) resolveFallback(call common.Call) *Registration[U] {
	if f.fallback.IsAbsent() {
		panic(newMissingImplementationError(f.info, call))
	}
	return f.fallback.GetOrPanic()
}

func (f *Fake[T, U]) consumeEphemeralCandidate(candidate *Registration[U]) {
	if candidate.IsEphemeral() {
		f.ephemerals = lo.Filter(f.ephemerals, func(item *Registration[U], _ int) bool {
			return item != candidate
		})
	}
}
