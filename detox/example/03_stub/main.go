package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// There are many types of test doubles:
	// - Dummies
	// - Spies
	// - Fakes
	// - Stubs
	// - Mocks
	// For us, a mock is simply a combination of all the others.

	// For dummies, its boils down to having a mock that can be instantiated
	// and passed around as a given interface.

	// For spies, we will cover this topic later on when talking about assertions.

	// For fakes, this is precisely what we configure when we are using Detox.

	// For stubs, due to some type system limitation, we are not able to isolate the
	// return type of mocked methods.
	// Thus, for Detox, a stub is basically a fake that would return hardcoded values.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	detox.When(mock.Detox, mock.Method3).Call(func() string {
		return "I am a stub"
	})
	// ACT
	fmt.Println(mock.Method3())
}
