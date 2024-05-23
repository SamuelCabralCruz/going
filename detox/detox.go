package detox

import (
	"fmt"
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

func (d *Detox) Reset() {
	d.fakes = map[string]any{}
	d.spies = map[string]any{}
}

func name[T any](mock *Detox, orig T) string {
	return fmt.Sprintf("%s.%s", mock.name, getId(orig))
}

func getId[T any](orig T) string {
	return phi.FunctionName(orig)
}

func resolveSpy[T any](mock *Detox, orig T) *spy.Spy {
	id := getId(orig)
	if s, ok := mock.spies[id].(*spy.Spy); ok {
		return s
	}
	newSpy := spy.NewSpy(mock.name, id)
	mock.spies[id] = newSpy
	return newSpy
}

func resolveFake[T any](mock *Detox, orig T) *fake.Fake[T] {
	id := getId(orig)
	if f, ok := mock.fakes[id].(*fake.Fake[T]); ok {
		return f
	}
	newFake := fake.NewFake[T](mock.name, id)
	mock.fakes[id] = newFake
	return newFake
}
