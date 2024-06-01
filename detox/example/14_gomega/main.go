//go:build example

package main

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example/00_boilerplate"
	. "github.com/SamuelCabralCruz/went/detox/matcher"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func main() {
	// Since we are strong advocates of testing with `ginkgo` and `gomega`, we
	// took the time to create dedicated gomega matchers for Botox.

	// Hence, each assertion method has its equivalent gomega matcher:
	// HasBeenCalled -> matcher.HaveBeenCalled
	// HasBeenCalledOnce -> matcher.HaveBeenCalledOnce
	// HasBeenCalledTimes -> matcher.HaveBeenCalledTimes
	// HasBeenCalledWith -> matcher.HaveBeenCalledWith
	// HasBeenCalledOnceWith -> matcher.HaveBeenCalledOnceWith
	// HasBeenCalledTimesWith -> matcher.HaveBeenCalledTimesWith
	// HasCalls -> matcher.HaveCalls
	// HasNthCall -> matcher.HaveNthCall
	// HasOrderedCalls -> matcher.HaveOrderedCalls

	// TEST
	// ARRANGE
	mock := example.NewSomeMock()
	mock.Default(example.SomeImplementation{})
	mocked := detox.When(mock.Detox, mock.Method5)
	// ACT
	mock.Method5("some text 1", true, []byte("ok-a"))
	mock.Method5("some text 2", false, []byte("ok-b"))
	mock.Method5("some text 3", true, []byte("ok-c"))
	// ASSERT
	It("should work", func() {
		Expect(mocked).To(HaveBeenCalled())
		Expect(mocked).NotTo(HaveBeenCalledOnce())
		Expect(mocked).To(HaveBeenCalledTimes(3))
		Expect(mocked).To(HaveBeenCalledWith("some text 2", false, []byte("ok-b")))
		Expect(mocked).To(HaveBeenCalledOnceWith("some text 3", true, []byte("ok-c")))
		Expect(mocked).NotTo(HaveBeenCalledTimesWith(2, "some text 1", true, []byte("ok-a")))
		Expect(mocked).To(HaveCalls(
			[]any{"some text 2", false, []byte("ok-b")},
			[]any{"some text 3", true, []byte("ok-c")},
			[]any{"some text 1", true, []byte("ok-a")},
		))
	})
}
