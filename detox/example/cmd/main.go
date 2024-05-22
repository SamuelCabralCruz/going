package main

import (
	"errors"
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example/pkg"
)

func fakedNamed(s string) (string, error) { return "ok", errors.New("coucou named function " + s) }

func main() {
	var myMock = pkg.SomeMockClass{Detox: detox.New[pkg.SomeMockClass]()}
	var myMock2 = pkg.SomeMockClass{Detox: detox.New[pkg.SomeMockClass]()}
	var myOtherMock = pkg.AnotherMockClass{Detox: detox.New[pkg.AnotherMockClass]()}

	// real impl
	detox.FakeImplementation(myMock.Detox, myMock.Hello, pkg.Impl{}.Hello)
	fmt.Println(myMock.Hello("after named fake"))

	// react to args
	detox.FakeImplementation(myMock.Detox, myMock.Hello, func(s string) (string, error) {
		if s == "sam" {
			return "goat", nil
		} else {
			return "lcm", nil
		}
	})

	detox.FakeImplementation(myMock.Detox, myMock.Hello, fakedNamed)
	fakedInline := func(s string) (string, error) { return "ok", errors.New("coucou inline function " + s) }
	detox.FakeImplementation(myMock2.Detox, myMock2.Hello, fakedInline)
	fmt.Println(myMock.Hello("after named fake"))
	fmt.Println(myMock.Hello("sam"))
	fmt.Println(myMock2.Hello("after named fake"))

	detox.FakeImplementation(myMock.Detox, myMock.Hello, fakedInline)

	fmt.Println(myMock.Hello("before method reset"))
	detox.Reset(myMock.Detox, myMock.Hello)
	//fmt.Println(myMock.Hello("after method reset")) // -> should fail

	detox.FakeImplementation(myMock.Detox, myMock.Hello, func(s string) (string, error) { return "ok", errors.New("coucou anonymous " + s) })

	fmt.Println(myMock.Hello("before mock reset"))
	myMock.Reset()
	//fmt.Println(myMock.Hello("after mock reset")) // -> should fail

	detox.FakeImplementationOnce(myMock.Detox, myMock.Hello, fakedNamed)
	detox.FakeImplementationOnce(myMock.Detox, myMock.Hello, fakedInline)
	detox.FakeImplementationOnce(myMock.Detox, myMock.Hello, func(s string) (string, error) { return "ok", errors.New("coucou anonymous " + s) })
	detox.FakeImplementation(myMock.Detox, myMock.Hello, func(s string) (string, error) { return "ok", errors.New("coucou remainder " + s) })

	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("2nd"))
	fmt.Println(myMock.Hello("3rd"))
	fmt.Println(myMock.Hello("4th"))
	fmt.Println(myMock.Hello("5th"))
	fmt.Println(myMock.Hello("6th"))

	detox.FakeImplementation(myMock.Detox, myMock.Prepare, func() pkg.Another {
		return myOtherMock
	})
	detox.FakeImplementation(myOtherMock.Detox, myOtherMock.Bye, func(s string) {
		fmt.Println("it works " + s)
	})
	myMock.Prepare().Bye("sam")

	fmt.Println(detox.Calls(myMock.Detox, myMock.Hello))
	fmt.Println(detox.CallsCount(myMock.Detox, myMock.Hello))
	fmt.Println(detox.NthCall(myMock.Detox, myMock.Hello, 3))

	myMock.Reset()

	fmt.Println(detox.Calls(myMock.Detox, myMock.Hello))
	fmt.Println(detox.CallsCount(myMock.Detox, myMock.Hello))
	//fmt.Println(detox.NthCall(myMock.Detox, myMock.Hello, 3)) // -> should fail

	// fluent api
	mocked := detox.When(myMock.Detox, myMock.Hello)
	mocked.CallOnce(func(s string) (string, error) {
		return "fluently faked 1", errors.New(s)
	})
	mocked.CallOnce(func(s string) (string, error) {
		return "fluently faked 2", errors.New(s)
	})
	mocked.Call(func(s string) (string, error) {
		return "fluently faked 3", errors.New(s)
	})
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("2nd"))
	fmt.Println(myMock.Hello("3rd"))
	fmt.Println(myMock.Hello("4th"))

	mocked.Reset()

	mocked.Call(func(s string) (string, error) {
		return "fluently persistent unconditional", errors.New(s)
	})
	mocked.CallOnce(func(s string) (string, error) {
		return "fluently ephemeral unconditional", errors.New(s)
	})
	// since this one is not ephemeral it can be called has many times as desired with 1st
	// should have priority over the previous one because it is specific to this use case
	mocked.WithArgs("1st").Call(func(s string) (string, error) {
		return "fluently conditional persistent", errors.New(s)
	})
	// will have priority over the previous one because it is ephemeral
	mocked.WithArgs("1st").CallOnce(func(s string) (string, error) {
		return "fluently conditional ephemeral", errors.New(s)
	})
	mocked.WithArgs("2nd").CallOnce(func(s string) (string, error) {
		return "fluently faked 2", errors.New(s)
	})
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("2nd"))
	//fmt.Println(myMock.Hello("3rd"))
	//fmt.Println(myMock.Hello("4th"))
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("1st"))
	fmt.Println(myMock.Hello("2nd"))

	// TODO: CallThrough (provide real)

	// TODO: create custom matchers
	// TODO: HaveBeenCalled() -> called at least once
	// TODO: HaveBeenCalledNth(int) -> called n times
	// TODO: HaveCalls([][]any) -> any order
	// TODO: HaveBeenCalledWith([]any) -> contains a calls with provided args
	// TODO: HaveCallSequence([][]any) -> specific order

	// TODO: caveats
	// TODO: Mock return values
}
