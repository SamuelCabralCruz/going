//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// We can mix and match all the registration types together. However, you need
	// to be aware of all they will interact.

	// Here are the relative priorities for fake implementation resolution from
	// the highest priority to the lowest:
	// 1. Ephemeral Conditional
	// 2. Persistent Conditional
	// 3. Ephemeral
	// 4. Persistent
	// 5. Default

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mock.Default(example.SomeImplementation{})
	mocked := detox.When(mock.Detox, mock.Method2)
	mocked.Call(func(arg string) {
		fmt.Println("persistent")
	})
	mocked.CallOnce(func(arg string) {
		fmt.Println("ephemeral")
	})
	mocked.WithArgs("conditional").Call(func(arg string) {
		fmt.Println("persistent conditional")
	})
	mocked.WithArgs("conditional").CallOnce(func(arg string) {
		fmt.Println("ephemeral conditional")
	})
	// ACT
	mock.Method2("conditional")          // -> should call ephemeral conditional
	mock.Method2("some input arg")       // -> should call ephemeral
	mock.Method1()                       // -> should call default
	mock.Method2("some other input arg") // -> should call persistent
	mock.Method3()                       // -> should call default
	mock.Method2("conditional")          // -> should call persistent conditional
	mock.Method2("conditional")          // -> should call persistent conditional
	mock.Method2("some other input arg") // -> should call persistent
	mock.Method2("conditional")          // -> should call persistent conditional
	mock.Method2("some other input arg") // -> should call persistent
	mock.Method2("conditional")          // -> should call persistent conditional
	mock.Method2("some other input arg") // -> should call persistent
}
