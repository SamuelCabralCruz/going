package usage_test

import (
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/detox/matcher"
	"github.com/SamuelCabralCruz/going/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(matcher.HaveBeenCalledOnce, func() {
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.NoArgNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
	})

	AfterEach(func() {
		mock.Reset()
	})

	Context("with called once mock", func() {
		BeforeEach(func() {
			mock.NoArgNoReturn()
		})

		It("should match", func() {
			Expect(mocked).To(matcher.HaveBeenCalledOnce())
		})
	})

	Context("with non called mock", func() {
		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveBeenCalledOnce())
		})
	})

	Context("with mock called more than once", func() {
		BeforeEach(func() {
			mock.NoArgNoReturn()
			mock.NoArgNoReturn()
		})

		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveBeenCalledOnce())
		})
	})
})
