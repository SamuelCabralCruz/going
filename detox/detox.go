package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/internal/spy"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/SamuelCabralCruz/went/phi"
)

func New[T any]() *Detox[T] {
	d := &Detox[T]{
		info: common.NewMockInfo[T](),
	}
	d.Reset()
	return d
}

type Detox[T any] struct {
	info        common.MockInfo
	defaultImpl optional.Optional[T]
	spies       map[string]any
	fakes       map[string]any
}

func (d *Detox[T]) Default(impl T) {
	d.defaultImpl = optional.Of(impl)
}

func (d *Detox[T]) Reset() {
	d.defaultImpl = optional.Empty[T]()
	d.fakes = map[string]any{}
	d.spies = map[string]any{}
}

func getId[T any](method T) string {
	return phi.FunctionName(method)
}

func resolveSpy[T any, U any](mock *Detox[T], method U) *spy.Spy {
	id := getId(method)
	if s, ok := mock.spies[id].(*spy.Spy); ok {
		return s
	}
	return createSpy(mock, method, id)
}

func createSpy[T any, U any](mock *Detox[T], method U, id string) *spy.Spy {
	newSpy := spy.NewSpy(common.NewMockedMethodInfo(mock.info, method))
	mock.spies[id] = newSpy
	return newSpy
}

func resolveFake[T any, U any](mock *Detox[T], method U) *fake.Fake[U] {
	id := getId(method)
	if f, ok := mock.fakes[id].(*fake.Fake[U]); ok {
		return f
	}
	return createFake(mock, method, id)
}

func createFake[T any, U any](mock *Detox[T], method U, id string) *fake.Fake[U] {
	newFake := fake.NewFake[U](common.NewMockedMethodInfo(mock.info, method), computeFallback[T, U](mock, id))
	mock.fakes[id] = newFake
	return newFake
}

func computeFallback[T any, U any](mock *Detox[T], id string) optional.Optional[U] {
	return optional.Transform(mock.defaultImpl, func(real T) U {
		candidate := phi.Value(real).MethodByName(id)
		return candidate.Interface().(U)
	})
}
