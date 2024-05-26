//go:build example

package main

import "github.com/SamuelCabralCruz/went/detox/example"

func main() {
	// To avoid your tests to panic or to partially mock the implementation, you
	// can provide a default implementation.

	// This default implementation will always be the last one to be resolved if
	// no other implementation is found.

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mock.Default(example.SomeImplementation{})
	// ACT
	mock.Method1()
	mock.Method2("arg2.1")
	_ = mock.Method3()
	_ = mock.Method4(4)
	mock.Method5("arg1", 2, true, []byte("four"))
	_ = mock.Method6('1', 'b', '3', 'd')
	_, _ = mock.Method7([]byte("first arg"), 2, "3", false)
}
