//go:build example

package main

import "github.com/SamuelCabralCruz/went/detox/example"

func main() {
	// While using your mocks, keep in mind that any invocation for which no
	// implementation has been provided will result in an error.

	// By error, we mean that your tests will panic. 😱

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	// ACT
	mock.Method1()
}
