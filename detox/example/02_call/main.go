//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// The simplest use case is to tell our mock to call a fake implementation.

	// This is what we call a persistent registration because this implementation
	// will always be resolved no matter the arguments and/or the number of times
	// it has been invoked.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	detox.When(mock.Detox, mock.Method2).Call(func(arg string) {
		fmt.Println(fmt.Sprintf("mocked persistent implementation - %+v", arg))
	})
	// ACT
	mock.Method2("1st")
	mock.Method2("2nd")
	mock.Method2("3rd")
	mock.Method2("4th")
}
