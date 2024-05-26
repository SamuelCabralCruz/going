package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// Another use case would be to fake an implementation a given number of times.

	// This is what we call an ephemeral registration because this implementation
	// will be forgotten at the moment it has been used.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	detox.When(mock.Detox, mock.Method2).CallOnce(func(arg string) {
		fmt.Println(fmt.Sprintf("mocked ephemeral implementation - %+v", arg))
	})
	// ACT
	mock.Method2("1st") // -> should consume the ephemeral
	mock.Method2("2nd") // -> should panic
}
