package fake

import (
	"github.com/SamuelCabralCruz/went/detox/internal/fake/faked"
	"github.com/samber/lo"
	"sort"
)

type priority int

const (
	ephemeralConditional priority = iota
	ephemeral
	persistentConditional
	persistent
)

type implementation[T any] struct {
	priority priority
	fake     faked.Faked[T]
}

func (i implementation[T]) isEphemeral() bool {
	return lo.Contains([]priority{ephemeralConditional, ephemeral}, i.priority)
}

func NewFake[T any](mockName string, methodName string) *Fake[T] {
	return &Fake[T]{
		mockName:   mockName,
		methodName: methodName,
	}
}

type Fake[T any] struct {
	mockName   string
	methodName string
	impl       []implementation[T]
}

func (f *Fake[T]) RegisterImplementation(impl T) {
	f.register(persistent, faked.NewUnconditionalFaked(impl))
}

func (f *Fake[T]) RegisterImplementationOnce(impl T) {
	f.register(ephemeral, faked.NewUnconditionalFaked(impl))
}

func (f *Fake[T]) RegisterConditionalImplementation(impl T, withArgs []any) {
	f.register(persistentConditional, faked.NewConditionalFaked(impl, withArgs))
}

func (f *Fake[T]) RegisterConditionalImplementationOnce(impl T, withArgs []any) {
	f.register(ephemeralConditional, faked.NewConditionalFaked(impl, withArgs))
}

func (f *Fake[T]) register(priority priority, fake faked.Faked[T]) {
	f.impl = append(f.impl, implementation[T]{priority, fake})
}

func (f *Fake[T]) InvokeFakeImplementation(args ...any) T {
	implementations := lo.Filter(f.impl, func(item implementation[T], _ int) bool {
		return item.fake.CanHandle(args...)
	})
	sort.Slice(implementations, func(i, j int) bool {
		return implementations[i].priority < implementations[j].priority
	})

	if len(implementations) == 0 {
		panic(newMissingFakeImplementationError(f.mockName, f.methodName))
	}

	impl := implementations[0]
	if impl.isEphemeral() {
		f.impl = lo.Filter(f.impl, func(item implementation[T], _ int) bool {
			return item != impl
		})
	}
	return impl.fake.Invoke()
}
