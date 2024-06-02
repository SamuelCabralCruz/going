//go:build test

package trust_test

import (
	"fmt"
	. "github.com/SamuelCabralCruz/went/gomicron/matcher"
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/trust"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(trust.AssertIsFunction, func() {
	var input any
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = trust.AssertIsFunction(input)
	}

	Context("with function as input", func() {
		BeforeEach(func() {
			input = func() { fmt.Println("i am a function") }
		})

		It("should return positive assertion", func() {
			act()

			Expect(observedValue).To(BeFunction(input))
			Expect(observedError).To(BeNil())
		})
	})

	Context("with non function as input", func() {
		BeforeEach(func() {
			input = "i am not a function"
		})

		It("should return negative assertion", func() {
			act()

			Expect(observedValue).To(BeNil())
			Expect(observedError).NotTo(BeNil())
			Expect(observedError.Error()).To(MatchRegexp(`Expected value to be a function, but was of type "string"`))
		})
	})
})

var _ = DescribeFunction(trust.IsFunction, func() {
	var input any
	var observedValue any
	var observedOk bool

	act := func() {
		observedValue, observedOk = trust.IsFunction(input)
	}

	Context("with function as input", func() {
		BeforeEach(func() {
			input = func() { fmt.Println("i am a function") }
		})

		It("should return positive validation", func() {
			act()

			Expect(observedValue).To(BeFunction(input))
			Expect(observedOk).To(BeTrue())
		})
	})

	Context("with non function as input", func() {
		BeforeEach(func() {
			input = "i am not a function"
		})

		It("should return negative validation", func() {
			act()

			Expect(observedValue).To(BeNil())
			Expect(observedOk).To(BeFalse())
		})
	})
})
