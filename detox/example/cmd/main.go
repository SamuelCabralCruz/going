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

	// TODO: fluent api
	//detox.When(myMock.Detox, myMock.Hello).Call(func() {})
	// TODO: WithArgs could simply be a filter
	// TODO: Call
	// TODO: CallOnce
	// TODO: CallThrough (provide real)
	// TODO: possibly Return --> doubt about that

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

	// TODO: create custom matchers
	// TODO: HaveBeenCalled() -> called at least once
	// TODO: HaveBeenCalledNth(int) -> called n times
	// TODO: HaveCalls([][]any) -> any order
	// TODO: HaveBeenCalledWith([]any) -> contains a calls with provided args
	// TODO: HaveCallSequence([][]any) -> specific order
}
