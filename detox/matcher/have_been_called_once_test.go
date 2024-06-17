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

var _ = DescribeFunction(matcher.HaveBeenCalledOnce, func() {
	var cut types.GomegaMatcher
	var actual any
	mock := fixture.NewInterface1Mock()
	mocked := detox.When(mock.Detox, mock.NoArgNoReturn)

	BeforeEach(func() {
		mock.Default(fixture.Implementation1{})
		actual = mocked
		cut = matcher.HaveBeenCalledOnce()
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
				It("should return false", func() {
					act()

					Expect(observedOk).To(BeFalse())
					Expect(observedError).To(BeNil())
				})
			})

			Context("with matching actual", func() {
				BeforeEach(func() {
					mock.NoArgNoReturn()
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

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(MatchRegexp("Expected Interface1\\.NoArgNoReturn \\(.*/matcher/have_been_called_once_test\\.go \\[18\\]\\) to have been called once, but was called 0 times"))
		})
	})

	DescribeFunction(types.GomegaMatcher.NegatedFailureMessage, func() {
		var observed string

		act := func() {
			observed = cut.NegatedFailureMessage(actual)
		}

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(MatchRegexp(
				"Expected Interface1\\.NoArgNoReturn \\(.*/matcher/have_been_called_once_test\\.go \\[18\\]\\) not to have been called once, but was called 0 times"))
		})
	})
})
