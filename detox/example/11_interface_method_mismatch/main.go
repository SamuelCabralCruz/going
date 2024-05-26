package main

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

type SomeOtherInterface interface {
	Method1() example.SomeInterface
}

func NewSomeOtherMock() SomeOtherMock {
	return SomeOtherMock{detox.New[SomeOtherInterface]()}
}

type SomeOtherMock struct {
	*detox.Detox[SomeOtherInterface]
}

var _ SomeOtherInterface = SomeOtherMock{}

func (s SomeOtherMock) Method1() example.SomeInterface {
	return detox.When(s.Detox, s.Method1).ResolveForArgs()()
}

func main() {
	// Based on the chaining example, some might have thought about trying to pass
	// a signature belonging to another mock with the same name but different type.

	// Detox will perform assertions to prevent such bad behavior.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	otherMock := NewSomeOtherMock()
	detox.When(mock.Detox, otherMock.Method1).Call(func() example.SomeInterface {
		return nil
	}) // -> should panic

	// Side Note:
	// Detox will also panic if you try to mock a method that does not exist on
	// the interface you are mocking.
	// If you are suspicious, you can always try it yourself. 😉
}
