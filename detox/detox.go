package detox

import (
	"github.com/SamuelCabralCruz/going/detox/internal/common"
	"github.com/SamuelCabralCruz/going/detox/internal/fake"
	"github.com/SamuelCabralCruz/going/detox/internal/spy"
	"github.com/SamuelCabralCruz/going/fn/optional"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/samber/lo"
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
	fakes       map[string]DefaultRegistrable[T]
}

type DefaultRegistrable[T any] interface {
	RegisterDefault(impl T)
}

func (d *Detox[T]) Default(impl T) {
	d.defaultImpl = optional.Of(impl)
	lo.ForEach(lo.Values(d.fakes), func(item DefaultRegistrable[T], _ int) {
		item.RegisterDefault(impl)
	})
}

func (d *Detox[T]) Reset() {
	d.defaultImpl = optional.Empty[T]()
	d.spies = map[string]any{}
	d.fakes = map[string]DefaultRegistrable[T]{}
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

func resolveFake[T any, U any](mock *Detox[T], method U) *fake.Fake[T, U] {
	id := getId(method)
	if f, ok := mock.fakes[id].(*fake.Fake[T, U]); ok {
		return f
	}
	return createFake(mock, method, id)
}

func createFake[T any, U any](mock *Detox[T], method U, id string) *fake.Fake[T, U] {
	newFake := fake.NewFake[T, U](common.NewMockedMethodInfo(mock.info, method), mock.defaultImpl)
	mock.fakes[id] = newFake
	return newFake
}
