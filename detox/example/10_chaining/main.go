package main

import (
	"fmt"
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
	// Another interesting use case would be the chaining of mocks. 🔗
	// What we mean by chain is the ability to make a mock return another mock.

	// Above, we implemented a new mock for another interface which has a method
	// that will return an implementation of `SomeInterface`.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	detox.When(mock.Detox, mock.Method1).Call(func() {
		fmt.Println("mock method1 invoked - chaining is working")
	})
	otherMock := NewSomeOtherMock()
	detox.When(otherMock.Detox, otherMock.Method1).Call(func() example.SomeInterface {
		fmt.Println("other mock method1 - return mock")
		return mock
	})

	// ACT
	otherMock.Method1().Method1()
}
