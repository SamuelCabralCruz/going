//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example"
)

func main() {
	// We like to keep the best for desert. 🫦

	// What would be a mock framework without any means to perform assertions?
	// With Detox, assertions are performed at the method level.
	// We are talking about the return of `detox.When` function, you got it right.

	// However, to avoid polluting the interface of a mocked method with all
	// the assertion methods, we exposed an `Assert` method to transform it
	// into an asserter.

	// Once transformed, you will then have access to multiple methods:
	// - HasBeenCalled
	// - HasBeenCalledOnce
	// - HasBeenCalledTimes
	// - HasBeenCalledWith
	// - HasBeenCalledOnceWith
	// - HasBeenCalledTimesWith
	// - HasCalls
	// - HasNthCall
	// - HasOrderedCalls

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mock.Default(example.SomeImplementation{})
	mocked := detox.When(mock.Detox, mock.Method5)
	mocked.Call(func(_ ...any) {})
	// ACT
	mock.Method5('1', '2', '3')
	mock.Method5('1')
	mock.Method5('2')
	mock.Method5('2')
	mock.Method5('2')
	mock.Method5('3')
	mock.Method5("some text", true, []byte("ok"))
	mock.Method1()
	// ASSERT
	fmt.Println("HasBeenCalled")
	fmt.Println(mocked.Assert().HasBeenCalled())                               // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method1).Assert().HasBeenCalled()) // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method2).Assert().HasBeenCalled()) // -> should be false
	fmt.Println("HasBeenCalledOnce")
	fmt.Println(detox.When(mock.Detox, mock.Method1).Assert().HasBeenCalledOnce()) // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method2).Assert().HasBeenCalledOnce()) // -> should be false
	fmt.Println(mocked.Assert().HasBeenCalledOnce())                               // -> should be false
	fmt.Println("HasBeenCalledTimes")
	fmt.Println(mocked.Assert().HasBeenCalledTimes(7))                               // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method1).Assert().HasBeenCalledTimes(1)) // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method2).Assert().HasBeenCalledTimes(0)) // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledTimes(6))                               // -> should be false
	fmt.Println(mocked.Assert().HasBeenCalledTimes(8))                               // -> should be false
	fmt.Println(mocked.Assert().HasBeenCalledTimes(-1))                              // -> should be false
	fmt.Println("HasBeenCalledWith")
	fmt.Println(mocked.Assert().HasBeenCalledWith('1'))                             // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledWith('2'))                             // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledWith("some text", true, []byte("ok"))) // -> should be true
	fmt.Println(detox.When(mock.Detox, mock.Method1).Assert().HasBeenCalledWith())  // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledWith('4'))                             // -> should be false
	fmt.Println("HasBeenCalledOnceWith")
	fmt.Println(mocked.Assert().HasBeenCalledOnceWith('1'))                             // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledOnceWith('3'))                             // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledOnceWith("some text", true, []byte("ok"))) // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledOnceWith('2'))                             // -> should be false
	fmt.Println(mocked.Assert().HasBeenCalledOnceWith("some", "other", "args"))         // -> should be false
	fmt.Println("HasBeenCalledTimesWith")
	fmt.Println(mocked.Assert().HasBeenCalledTimesWith(1, '1', '2', '3')) // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledTimesWith(3, '2'))           // -> should be true
	fmt.Println(mocked.Assert().HasBeenCalledTimesWith(4, '2'))           // -> should be false
	fmt.Println(mocked.Assert().HasBeenCalledTimesWith(2, '2'))           // -> should be false
	fmt.Println("HasCalls")
	fmt.Println(detox.When(mock.Detox, mock.Method2).Assert().HasCalls())                                     // -> should be true
	fmt.Println(mocked.Assert().HasCalls())                                                                   // -> should be true
	fmt.Println(mocked.Assert().HasCalls(detox.Call{'1', '2', '3'}))                                          // -> should be true
	fmt.Println(mocked.Assert().HasCalls(detox.Call{'1'}, detox.Call{'3'}, detox.Call{'2'}))                  // -> should be true
	fmt.Println(mocked.Assert().HasCalls(detox.Call{'1'}, detox.Call{'2'}))                                   // -> should be true
	fmt.Println(mocked.Assert().HasCalls(detox.Call{'1'}, detox.Call{'4'}, detox.Call{'3'}, detox.Call{'2'})) // -> should be false
	fmt.Println(mocked.Assert().HasCalls(detox.Call{'1'}, detox.Call{'1'}))                                   // -> should be false
	fmt.Println("HasNthCall")
	fmt.Println(mocked.Assert().HasNthCall(0, detox.Call{'1', '2', '3'})) // -> should be true
	fmt.Println(mocked.Assert().HasNthCall(1, detox.Call{'1'}))           // -> should be true
	fmt.Println(mocked.Assert().HasNthCall(4, detox.Call{'2'}))           // -> should be true
	fmt.Println(mocked.Assert().HasNthCall(5, detox.Call{'3'}))           // -> should be true
	fmt.Println(mocked.Assert().HasNthCall(4, detox.Call{'3'}))           // -> should be false
	fmt.Println(mocked.Assert().HasNthCall(6, detox.Call{'3'}))           // -> should be false
	fmt.Println("HasOrderedCalls")
	fmt.Println(mocked.Assert().HasOrderedCalls(
		detox.Call{'1', '2', '3'},
		detox.Call{'1'},
		detox.Call{'2'},
		detox.Call{'2'},
		detox.Call{'2'},
		detox.Call{'3'},
		detox.Call{"some text", true, []byte("ok")},
	)) // -> should be true
	fmt.Println(mocked.Assert().HasOrderedCalls(
		detox.Call{'1', '2', '3'},
		detox.Call{'1'},
	)) // -> should be false
	fmt.Println(mocked.Assert().HasOrderedCalls(
		detox.Call{'1', '2', '3'},
		detox.Call{'2'},
		detox.Call{'1'},
		detox.Call{'2'},
		detox.Call{'2'},
		detox.Call{'3'},
		detox.Call{"some text", true, []byte("ok")},
	)) // -> should be false
}
