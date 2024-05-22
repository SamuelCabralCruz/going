package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/internal/spy"
	"github.com/SamuelCabralCruz/went/phi"
)

func New[T any]() *Detox {
	d := &Detox{name: phi.BaseTypeName[T]()}
	d.Reset()
	return d
}

type Detox struct {
	name  string
	spies map[string]any
	fakes map[string]any
}

func (d *Detox) Name() string {
	return d.name
}

func (d *Detox) Reset() {
	d.fakes = map[string]any{}
	d.spies = map[string]any{}
}

func getId[T any](orig T) string {
	return phi.FunctionName(orig)
}

func resolveSpy[T any](mock *Detox, orig T) *spy.Spy {
	id := getId(orig)
	if spy, ok := mock.spies[id].(*spy.Spy); ok {
		return spy
	}
	newSpy := spy.NewSpy(mock.name, id)
	mock.spies[id] = newSpy
	return newSpy
}

func resolveFake[T any](mock *Detox, orig T) *fake.Fake[T] {
	id := getId(orig)
	if fake, ok := mock.fakes[id].(*fake.Fake[T]); ok {
		return fake
	}
	newFake := fake.NewFake[T](mock.name, id)
	mock.fakes[id] = newFake
	return newFake
}

func FakeImplementation[T any](mock *Detox, orig T, impl T) {
	resolveFake(mock, orig).RegisterImplementation(impl)
}

func FakeImplementationOnce[T any](mock *Detox, orig T, impl T) {
	resolveFake(mock, orig).RegisterImplementationOnce(impl)
}

func FakeConditionalImplementation[T any](mock *Detox, orig T, impl T, withArgs []any) {
	resolveFake(mock, orig).RegisterConditionalImplementation(impl, withArgs)
}

func FakeConditionalImplementationOnce[T any](mock *Detox, orig T, impl T, withArgs []any) {
	resolveFake(mock, orig).RegisterConditionalImplementationOnce(impl, withArgs)
}

func Reset[T any](mock *Detox, orig T) {
	id := getId(orig)
	delete(mock.fakes, id)
	delete(mock.spies, id)
}

func RegisterInvocation[T any](mock *Detox, orig T, args ...any) {
	resolveSpy(mock, orig).RegisterInvocation(args...)
}

func InvokeFakeImplementation[T any](mock *Detox, orig T, args ...any) T {
	return resolveFake(mock, orig).InvokeFakeImplementation(args...)
}

func Calls[T any](mock *Detox, orig T) [][]any {
	return resolveSpy(mock, orig).Calls()
}

func CallsCount[T any](mock *Detox, orig T) int {
	return resolveSpy(mock, orig).CallsCount()
}

func NthCall[T any](mock *Detox, orig T, index int) []any {
	return resolveSpy(mock, orig).NthCall(index)
}
