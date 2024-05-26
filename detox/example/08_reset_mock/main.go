//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// Another important feature of a mock is being able to reset it to avoid
	// side effect between tests.

	// To reset a mock, we can simply invoke the `Reset` method on the mock itself.
	// This will forget ALL fakes and metrics accumulated by the mock.

	// TEST 1
	// ARRANGE
	mock := example.NewSomeMock()
	mock.Default(example.SomeImplementation{})
	detox.When(mock.Detox, mock.Method1).Call(func() {
		fmt.Println("method1 fake implementation")
	})
	// ACT
	mock.Method1()
	mock.Method2("something")

	// AFTER EACH
	mock.Reset()

	// TEST 2
	// ARRANGE
	// ACT
	mock.Method1() // -> should panic
}
