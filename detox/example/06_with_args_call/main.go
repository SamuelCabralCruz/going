//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example/00_boilerplate"
)

func main() {
	// Using Detox, we can also fake an implementation conditionally to the input
	// arguments.

	// This feature can be particularly useful when we want our mock to react
	// differently based on the provided arguments.
	// It can also be used to implicitly validate which arguments have been provided
	// to our mock instead of doing after the fact.

	// This is what we call a persistent conditional registration because this
	// implementation will always be resolved no matter the number of times
	// it has been invoked as long as the input arguments match the configured ones.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mocked := detox.When(mock.Detox, mock.Method2)
	mocked.WithArgs("1st").Call(func(arg string) {
		fmt.Println("mocked persistent conditional implementation - A")
	})
	mocked.WithArgs("2nd").Call(func(arg string) {
		fmt.Println("mocked persistent conditional implementation - B")
	})
	// ACT
	mock.Method2("1st")
	mock.Method2("2nd")
	mock.Method2("2nd")
	mock.Method2("1st")
	mock.Method2("2nd")
	mock.Method2("1st")
	mock.Method2("1st")
	mock.Method2("2nd")
	mock.Method2("1st")
	mock.Method2("3rd") // -> should panic
}
