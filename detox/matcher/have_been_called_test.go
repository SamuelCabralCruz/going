//go:build test

package matcher_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/matcher"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.NoArgNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
		mocked.Call(func() {})
	})

	AfterEach(func() {
		mock.Reset()
	})

	DescribeFunction(matcher.HaveBeenCalled, func() {
		It("should determine whether mock has been called", func() {
			mock.NoArgNoReturn()

			Expect(mocked).To(matcher.HaveBeenCalled())
		})
	})
})
