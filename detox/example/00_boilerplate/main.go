//go:build example

package example

import (
	"errors"
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
)

// Detox has been designed on the principle that we should only use mock when
// there is a contract/interface in place.

// Side note:
// We consider the creation of an interface for the sole purpose of
// mocking a class as a smell which might indicate flaws in your design.

// SomeInterface
// This is the starting point: an interface.
type SomeInterface interface {
	Method1()
	Method2(string)
	Method3() string
	Method4(int) bool
	Method5(...any)
	Method6(...byte) []byte
	Method7([]byte, ...any) (bool, error)
}

// SomeImplementation
// Then, you will normally have some implementation of this interface in your
// production code.
type SomeImplementation struct {
}

// Here, we make sure that `SomeImplementation` will keep inline with `SomeInterface`.
// We recommend always using this trick in your own code to force the compiler to
// fail early in case of an implementation drift.
var _ SomeInterface = SomeImplementation{}

// The following is simply our production code.

func (s SomeImplementation) Method1() {
	fmt.Println("Production code - Method 1")
}

func (s SomeImplementation) Method2(arg string) {
	fmt.Println(fmt.Sprintf("Production code - Method 2 - %+v", arg))
}

func (s SomeImplementation) Method3() string {
	fmt.Println("Production code - Method 3")
	return "some value"
}

func (s SomeImplementation) Method4(arg int) bool {
	fmt.Println(fmt.Sprintf("Production code - Method 4 - %+v", arg))
	return true
}

func (s SomeImplementation) Method5(args ...any) {
	fmt.Println(fmt.Sprintf("Production code - Method 6 - %+v", args))
}

func (s SomeImplementation) Method6(args ...byte) []byte {
	fmt.Println(fmt.Sprintf("Production code - Method 6 - %+v", args))
	return []byte("some value")
}

func (s SomeImplementation) Method7(arg1 []byte, arg2 ...any) (bool, error) {
	fmt.Println(fmt.Sprintf("Production code - Method 7 - %+v", append([]any{arg1}, arg2...)))
	return false, errors.New("some error message")
}

// Until now, there is nothing new brought by Detox. All the stuff above represents
// parts of your own application that you want to mock.
// This is where the magic happens.

// SomeMock
// We first create a new struct that will embed a reference on a Detox configured
// for the interface we want to mock (SomeInterface).
type SomeMock struct {
	*detox.Detox[SomeInterface]
}

// NewSomeMock
// We recommend creating a constructor for your mocks to simplify there instantiation.
// That way, you can be sure that the mock will stay on track with its associated interface.
func NewSomeMock() SomeMock {
	return SomeMock{detox.New[SomeInterface]()}
}

// Then, we once again enforce the implementation of SomeInterface by SomeMock.
var _ SomeInterface = SomeMock{}

// ⚠️ This part is important! ⚠️
// This could feel like some heavy lifting, we know...
// This is the best we could achieve.

// We might consider a code generator to avoid this step in the future.
// On the other hand, avoiding code generation is the reason why we started this
// initiative.
// We are not fans of code generation for the following reasons:
// - generated code is usually hard to understand and highly verbose
// - code generator usually offers no path forward in case of breaking changes
// - generated mock classes are often denatured from the original interface
// - we want to allow the user to edit their mocks without fear of breaking underlying
//   behaviors
// - we don't want to pollute our project with mock generation concerns

// Anyway, if you are still reading, it means that you at least agree a little
// with our concerns.

// So, to create the mock class, you will need to do the following for each method
// of the interface:
// - Use the `When` function exposed by the detox package to create a mocked method instance
// - Pass in the Detox reference and the method you are currently mocking
// - On the returned mocked method instance, call `ResolveForArgs` by forwarding the arguments
// - Invoke the resolved implementation with the arguments
// - Return the result of the invocation if the method is supposed to return anything

func (s SomeMock) Method1() {
	detox.When(s.Detox, s.Method1).ResolveForArgs()()
}

func (s SomeMock) Method2(arg string) {
	detox.When(s.Detox, s.Method2).ResolveForArgs(arg)(arg)
}

func (s SomeMock) Method3() string {
	return detox.When(s.Detox, s.Method3).ResolveForArgs()()
}

func (s SomeMock) Method4(arg int) bool {
	return detox.When(s.Detox, s.Method4).ResolveForArgs(arg)(arg)
}

func (s SomeMock) Method5(args ...any) {
	detox.When(s.Detox, s.Method5).ResolveForArgs(args...)(args...)
}

func (s SomeMock) Method6(args ...byte) []byte {
	// Here, we see an edge case where we need to convert args from []byte to
	// []any before forwarding them to the `ResolveForArgs` method.
	forwarded := make([]any, len(args))
	for i, a := range args {
		forwarded[i] = a
	}
	return detox.When(s.Detox, s.Method6).ResolveForArgs(forwarded...)(args...)
}

func (s SomeMock) Method7(arg1 []byte, arg2 ...any) (bool, error) {
	// Here, we see an edge case where we need to combine args before forwarding
	// them to the `ResolveForArgs` method because of the variadic behavior
	// alongside other arguments.
	args := append([]any{arg1}, arg2...)
	return detox.When(s.Detox, s.Method7).ResolveForArgs(args...)(arg1, arg2...)
}

// That's it! 🎉
// You now have a fully working mock implementation of your contract.
// You can now start using it in your tests.

// Be honest, it was a piece of 🧁.
