package usage_test

import (
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/detox/matcher"
	"github.com/SamuelCabralCruz/going/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(matcher.HaveBeenCalledTimesWith, func() {
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.MultipleArgsNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
	})

	AfterEach(func() {
		mock.Reset()
	})

	Context("with mock having desired calls", func() {
		BeforeEach(func() {
			mock.MultipleArgsNoReturn(2, false, []byte{'b', 'a', 'c'})
			mock.MultipleArgsNoReturn(1, true, []byte{'a', 'b', 'c'})
			mock.MultipleArgsNoReturn(2, false, []byte{'b', 'a', 'c'})
			mock.MultipleArgsNoReturn(3, true, []byte{'c', 'b', 'a'})
			mock.MultipleArgsNoReturn(2, false, []byte{'b', 'a', 'c'})
		})

		It("should match", func() {
			Expect(mocked).To(matcher.HaveBeenCalledTimesWith(3, 2, false, []byte{'b', 'a', 'c'}))
		})
	})

	Context("with mock not having desired calls", func() {
		BeforeEach(func() {
			mock.MultipleArgsNoReturn(1, true, []byte{'a', 'b', 'c'})
			mock.MultipleArgsNoReturn(2, false, []byte{'b', 'a', 'c'})
			mock.MultipleArgsNoReturn(3, true, []byte{'c', 'b', 'a'})
		})

		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveBeenCalledTimesWith(3, 1, false, []byte{'c', 'b', 'a'}))
		})
	})
})
