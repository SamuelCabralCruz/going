package scallop_test

import (
	"fmt"
	. "github.com/SamuelCabralCruz/going/kinggo"
	"github.com/SamuelCabralCruz/going/scallop"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(scallop.AsAnySlice[any], func() {
	var cut []string
	var observed []any

	act := func() {
		observed = scallop.AsAnySlice(cut)
	}

	BeforeEach(func() {
		cut = []string{"a", "b", "c"}
	})

	It("should return the inputted slice properly typed", func() {
		act()

		Expect(observed).To(HaveLen(len(cut)))
		Expect(fmt.Sprintf("%+v", observed)).To(Equal(fmt.Sprintf("%+v", cut)))
		Expect(observed).To(BeAssignableToTypeOf([]any{}))
	})
})

var _ = DescribeFunction(scallop.Copy[any], func() {
	var cut []string
	var observed []string

	act := func() {
		observed = scallop.Copy(cut)
	}

	BeforeEach(func() {
		cut = []string{"a", "b", "c"}
	})

	It("should make a shallow copy of the inputted slice", func() {
		act()

		Expect(observed).NotTo(BeIdenticalTo(cut))
		Expect(observed).To(BeEquivalentTo(cut))
	})
})

var _ = DescribeFunction(scallop.Pop[any], func() {
	var cut []string
	var observedValue string
	var observedRemainingValues []string
	var observedError error

	act := func() {
		observedValue, observedRemainingValues, observedError = scallop.Pop(cut)
	}

	Context("with empty slice", func() {
		BeforeEach(func() {
			cut = []string{}
		})

		It("should return an error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedRemainingValues).To(BeZero())
			Expect(observedError).NotTo(BeNil())
			Expect(observedError).To(BeAssignableToTypeOf(scallop.IndexOutOfBoundsError{}))
		})
	})

	Context("with non empty slice", func() {
		BeforeEach(func() {
			cut = []string{"a", "b", "c"}
		})

		It("should return first element and remaining slice", func() {
			act()

			Expect(observedValue).To(Equal("a"))
			Expect(observedRemainingValues).To(Equal([]string{"b", "c"}))
			Expect(observedError).To(BeNil())
		})
	})
})
