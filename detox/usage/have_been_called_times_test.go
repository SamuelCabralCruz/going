package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/matcher"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(matcher.HaveBeenCalledTimes, func() {
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.NoArgNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
	})

	AfterEach(func() {
		mock.Reset()
	})

	Context("with mock called exact number of times", func() {
		BeforeEach(func() {
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
		})

		It("should match", func() {
			Expect(mocked).To(matcher.HaveBeenCalledTimes(3))
		})
	})

	Context("with mock called one too many times", func() {
		BeforeEach(func() {
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
		})

		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveBeenCalledTimes(3))
		})
	})

	Context("with mock called one too few times", func() {
		BeforeEach(func() {
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
		})

		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveBeenCalledTimes(3))
		})
	})
})
