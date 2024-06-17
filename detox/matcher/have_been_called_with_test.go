//go:build test

package matcher_test

import (
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/detox/matcher"
	"github.com/SamuelCabralCruz/going/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = DescribeFunction(matcher.HaveBeenCalledWith, func() {
	var cut types.GomegaMatcher
	var inputs []any
	var actual any
	expectedArg := "some input arg"
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.SingleArgNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
		actual = mocked
		inputs = []any{expectedArg}
		cut = matcher.HaveBeenCalledWith(inputs...)
	})

	AfterEach(func() {
		mock.Reset()
	})

	DescribeFunction(types.GomegaMatcher.Match, func() {
		var observedOk bool
		var observedError error

		act := func() {
			observedOk, observedError = cut.Match(actual)
		}

		Context("with invalid actual", func() {
			BeforeEach(func() {
				actual = 1234
			})

			It("should return error", func() {
				act()

				Expect(observedOk).To(BeFalse())
				Expect(observedError).NotTo(BeNil())
				Expect(observedError.Error()).To(ContainSubstring("Expected \"1234\" to be of type \"Assertable\", but was of type \"int\""))
			})
		})

		Context("with valid actual", func() {
			Context("with non matching actual", func() {
				BeforeEach(func() {
					mock.SingleArgNoReturn("first")
					mock.SingleArgNoReturn("second")
				})

				It("should return false", func() {
					act()

					Expect(observedOk).To(BeFalse())
					Expect(observedError).To(BeNil())
				})
			})

			Context("with matching actual", func() {
				BeforeEach(func() {
					mock.SingleArgNoReturn("first")
					mock.SingleArgNoReturn(expectedArg)
					mock.SingleArgNoReturn("second")
				})

				It("should return true", func() {
					act()

					Expect(observedOk).To(BeTrue())
					Expect(observedError).To(BeNil())
				})
			})
		})
	})

	DescribeFunction(types.GomegaMatcher.FailureMessage, func() {
		var observed string

		act := func() {
			observed = cut.FailureMessage(actual)
		}

		BeforeEach(func() {
			mock.SingleArgNoReturn("first")
			mock.SingleArgNoReturn("second")
		})

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(MatchRegexp("Expected Interface1\\.SingleArgNoReturn \\(.*/matcher/have_been_called_with_test\\.go \\[20\\]\\) to have been called at least once with following args:\\n\\t\\t\\[<string> some input arg\\]\\n, but received calls were:\\n\\t\\t\\[0\\]: \\[<string> first\\]\\n\\t\\t\\[1\\]: \\[<string> second\\]"))
		})
	})

	DescribeFunction(types.GomegaMatcher.NegatedFailureMessage, func() {
		var observed string

		act := func() {
			observed = cut.NegatedFailureMessage(actual)
		}

		BeforeEach(func() {
			mock.SingleArgNoReturn("first")
			mock.SingleArgNoReturn("second")
		})

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(MatchRegexp("Expected Interface1\\.SingleArgNoReturn \\(.*/matcher/have_been_called_with_test\\.go \\[20\\]\\) not to have been called at least once with following args:\\n\\t\\t\\[<string> some input arg\\]\\n, but received calls were:\\n\\t\\t\\[0\\]: \\[<string> first\\]\\n\\t\\t\\[1\\]: \\[<string> second\\]"))
		})
	})
})
