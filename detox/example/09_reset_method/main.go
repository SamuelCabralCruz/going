//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// Alternatively, you can use the `Reset` method on the return of the
	// `detox.When` function.

	// This would have the effect of forgetting all fakes and spies for a specific
	// method.

	// TEST 1
	// ARRANGE
	mock := example.NewSomeMock()
	mocked1 := detox.When(mock.Detox, mock.Method1)
	mocked1.Call(func() {
		fmt.Println("method1 persistent fake implementation")
	})
	mocked1.CallOnce(func() {
		fmt.Println("method1 ephemeral fake implementation")
	})
	mocked2 := detox.When(mock.Detox, mock.Method2)
	mocked2.Call(func(arg string) {
		fmt.Println(fmt.Sprintf("method2 persistent fake implementation - %+v", arg))
	})
	// ACT
	mock.Method1()
	mock.Method2("something")
	mock.Method2("something else")
	mock.Method1()
	mock.Method1()

	// AFTER EACH
	mocked1.Reset()

	// TEST 2
	// ARRANGE
	// ACT
	mock.Method2("something")
	mock.Method2("something else")
	mock.Method1() // -> should panic
}
