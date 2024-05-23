package detox_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example/pkg"
	. "github.com/SamuelCabralCruz/went/detox/matcher"
	. "github.com/SamuelCabralCruz/went/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox](func() {
	mock := pkg.NewSomeMockClass()
	mockedHello := detox.When(mock.Detox, mock.Hello)

	BeforeEach(func() {
		mockedHello.Call(pkg.Impl{}.Hello)
	})

	AfterEach(func() {
		mock.Reset()
	})

	DescribeFunction(HaveBeenCalled, func() {
		It("should determine whether mock has been called", func() {
			_, _ = mock.Hello("some arg")

			Expect(mockedHello).To(HaveBeenCalled())
		})
	})
})
