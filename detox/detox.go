package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal"
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

func getId[T any](orig T) string {
	return phi.FunctionName(orig)
}

func resolveSpy[T any](mock *Detox, orig T) *internal.Spy {
	id := getId(orig)
	if spy, ok := mock.spies[id].(*internal.Spy); ok {
		return spy
	}
	newSpy := internal.NewSpy()
	mock.spies[id] = newSpy
	return newSpy
}

func resolveFake[T any](mock *Detox, orig T) *internal.Fake[T] {
	id := getId(orig)
	if fake, ok := mock.fakes[id].(*internal.Fake[T]); ok {
		return fake
	}
	newFake := internal.NewFake[T](mock.name, id)
	mock.fakes[id] = newFake
	return newFake
}

func FakeImplementation[T any](mock *Detox, orig T, impl T) {
	resolveFake(mock, orig).RegisterImplementation(impl)
}

func FakeImplementationOnce[T any](mock *Detox, orig T, impl T) {
	resolveFake(mock, orig).RegisterImplementationOnce(impl)
}

func (d *Detox) Reset() {
	d.fakes = map[string]any{}
	d.spies = map[string]any{}
}

func Reset[T any](mock *Detox, orig T) {
	id := getId(orig)
	delete(mock.fakes, id)
	delete(mock.spies, id)
}

func RegisterInvocation[T any](mock *Detox, orig T, args ...any) {
	resolveSpy(mock, orig).RegisterInvocation(args...)
}

func InvokeFakeImplementation[T any](mock *Detox, orig T) T {
	return resolveFake(mock, orig).InvokeFakeImplementation()
}

func Calls[T any](mock *Detox, orig T) [][]any {
	return resolveSpy(mock, orig).Calls()
}

func CallsCount[T any](mock *Detox, orig T) int {
	return len(Calls(mock, orig))
}

func NthCall[T any](mock *Detox, orig T, index int) []any {
	count := CallsCount(mock, orig)
	if index >= count {
		panic(newInvalidCallIndexError(mock.name, index, count))
	}
	return Calls(mock, orig)[index]
}
