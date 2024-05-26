package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// Similarly, you can also use the `CallOnce` method on the return of `WithArgs`.

	// I bet you already figured out that we would call the following an ephemeral
	// conditional registration. Big Brain 🧠, I must admit.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mocked := detox.When(mock.Detox, mock.Method2)
	mocked.WithArgs("1st").CallOnce(func(arg string) {
		fmt.Println("mocked ephemeral conditional implementation - A.1")
	})
	mocked.WithArgs("1st").CallOnce(func(arg string) {
		fmt.Println("mocked ephemeral conditional implementation - A.2")
	})
	mocked.WithArgs("2nd").CallOnce(func(arg string) {
		fmt.Println("mocked ephemeral conditional implementation - B")
	})
	// ACT
	mock.Method2("1st") // -> should consume ephemeral (A.1)
	mock.Method2("2nd") // -> should consume ephemeral (B)
	mock.Method2("1st") // -> should consume ephemeral (A.2)
	mock.Method2("1st") // -> should panic

	// Side Note:
	// Observe that ephemeral registrations are resolved in the order they have
	// been defined.
	// First In, First Out.
}
