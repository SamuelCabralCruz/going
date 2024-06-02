package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/matcher"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(matcher.HaveNthCall, func() {
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
			mock.MultipleArgsNoReturn(5, false, []byte{'d', 'e', 'f'})
		})

		It("should match", func() {
			Expect(mocked).To(matcher.HaveNthCall(4, []any{5, false, []byte{'d', 'e', 'f'}}))
		})
	})

	Context("with mock not having desired calls", func() {
		BeforeEach(func() {
			mock.MultipleArgsNoReturn(1, true, []byte{'a', 'b', 'c'})
			mock.MultipleArgsNoReturn(2, false, []byte{'b', 'a', 'c'})
			mock.MultipleArgsNoReturn(3, true, []byte{'c', 'b', 'a'})
		})

		It("should not match", func() {
			Expect(mocked).NotTo(matcher.HaveNthCall(1, []any{5, false, []byte{'d', 'e', 'f'}}))
		})
	})

	Context("with expected index out of bounds", func() {
		It("should not match", func() {
			Expect(mocked).To(matcher.HaveNthCall(10, []any{}))
		})
	})
})
