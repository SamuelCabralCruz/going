package matcher_test

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/gomicron/matcher"
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/xpctd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var validInput = func() { fmt.Println("some function") }
var invalidInput = 1234
var matchingActual = validInput
var nonMatchingActual = func() { fmt.Println("some other function") }
var invalidActual = invalidInput

var _ = DescribeFunction(matcher.BeFunction, func() {
	var cut types.GomegaMatcher
	var input any
	var actual any

	Describe("initialization", func() {
		act := func() {
			matcher.BeFunction(input)
		}

		Context("with invalid input", func() {
			BeforeEach(func() {
				input = invalidInput
			})

			It("should panic", func() {
				Expect(act).To(PanicWith(BeAssignableToTypeOf(xpctd.ExpectationError{})))
			})
		})

		Context("with function as input", func() {
			BeforeEach(func() {
				input = validInput
			})

			It("should not panic", func() {
				Expect(act).NotTo(Panic())
			})
		})
	})

	BeforeEach(func() {
		input = validInput
		cut = matcher.BeFunction(input)
	})

	DescribeFunction(types.GomegaMatcher.Match, func() {
		var observedOk bool
		var observedError error

		act := func() {
			observedOk, observedError = cut.Match(actual)
		}

		Context("with non function as actual", func() {
			BeforeEach(func() {
				actual = invalidActual
			})

			It("should return error", func() {
				act()

				Expect(observedOk).To(BeFalse())
				Expect(observedError).NotTo(BeNil())
				Expect(observedError.Error()).To(ContainSubstring("Expected value to be a function, but was of type \"int\""))
			})
		})

		Context("with function as actual", func() {
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

			Expect(observed).To(MatchRegexp(
				"Expected .*/matcher_test.init.func2 to be identical to .*/matcher_test.init.func1"))
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

			Expect(observed).To(MatchRegexp(
				"Expected .*/matcher_test.init.func2 not to be identical to .*/matcher_test.init.func1"))
		})
	})
})
