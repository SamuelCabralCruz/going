//go:build test

package fn_test

import (
	"github.com/SamuelCabralCruz/went/fn"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(fn.Identity[any], func() {
	var input any
	var observed any

	act := func() {
		observed = fn.Identity(input)
	}

	BeforeEach(func() {
		input = "some value"
	})

	It("should return inputted value", func() {
		act()

		Expect(observed).To(Equal(input))
	})
})
