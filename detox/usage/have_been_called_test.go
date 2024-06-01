package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/matcher"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(matcher.HaveBeenCalled, func() {
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
		Context("with called mock", func() {
			BeforeEach(func() {
				mock.NoArgNoReturn()
			})

			It("should match", func() {
				Expect(mocked).To(matcher.HaveBeenCalled())
			})
		})

		Context("with non called mock", func() {
			It("should not match", func() {
				Expect(mocked).NotTo(matcher.HaveBeenCalled())
			})
		})
	})
})
