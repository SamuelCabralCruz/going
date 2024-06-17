//go:build test

package tuple_test

import (
	"github.com/SamuelCabralCruz/going/fn/tuple"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(tuple.Swap[any, any], func() {
	var inputLeft any
	var inputRight any
	var observedLeft any
	var observedRight any

	act := func() {
		observedLeft, observedRight = tuple.Swap(inputLeft, inputRight)
	}

	BeforeEach(func() {
		inputLeft = "left"
		inputRight = "right"
	})

	It("should swap inputs", func() {
		act()

		Expect(observedLeft).To(Equal(inputRight))
		Expect(observedRight).To(Equal(inputLeft))
	})
})
