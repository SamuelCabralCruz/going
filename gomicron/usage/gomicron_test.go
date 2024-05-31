//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/gomicron"
	"github.com/SamuelCabralCruz/went/gomicron/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/xpctd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

const matchingActual = "hohoho"
const nonMatchingActual = "some invalid value"
const invalidActual = 1234

var _ = DescribeFunction(gomicron.ToGomegaMatcher[any], func() {
	var cut types.GomegaMatcher
	var actual any

	BeforeEach(func() {
		cut = fixture.BeSomeCustomMatcher()
		actual = nil
	})

	DescribeFunction(types.GomegaMatcher.Match, func() {
		var observedOk bool
		var observedError error

		act := func() {
			observedOk, observedError = cut.Match(actual)
		}

		Context("with actual of undesired type", func() {
			BeforeEach(func() {
				actual = invalidActual
			})

			It("should return error", func() {
				act()

				Expect(observedOk).To(BeFalse())
				Expect(observedError).To(BeAssignableToTypeOf(xpctd.ExpectationError{}))
			})
		})

		Context("with actual of desired type", func() {
			Context("with non matching actual", func() {
				BeforeEach(func() {
					actual = nonMatchingActual
				})

				It("should return false", func() {
					act()

					Expect(observedOk).To(BeFalse())
					Expect(observedError).To(BeNil())
				})
			})

			Context("with matching actual", func() {
				BeforeEach(func() {
					actual = matchingActual
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
			actual = nonMatchingActual
		})

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(Equal("Expected \"some invalid value\" to match"))
		})
	})

	DescribeFunction(types.GomegaMatcher.NegatedFailureMessage, func() {
		var observed string

		act := func() {
			observed = cut.NegatedFailureMessage(actual)
		}

		BeforeEach(func() {
			actual = nonMatchingActual
		})

		It("should return failure message properly formatted", func() {
			act()

			Expect(observed).To(Equal("Expected \"some invalid value\" not to match"))
		})
	})
})
