//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/usage"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	DescribeFunction(detox.When[any, any], func() {
		var cut usage.Interface2Mock
		var method any

		act := func() {
			detox.When(cut.Detox, method)
		}

		BeforeEach(func() {
			cut = usage.NewInterface2Mock()
		})

		Context("with method name not belonging to mocked interface", func() {
			BeforeEach(func() {
				method = usage.NewInterface3Mock().Method
			})

			It("should panic", func() {
				Expect(act).To(PanicWith(BeAssignableToTypeOf(detox.InterfaceMethodMismatchError{})))
			})
		})

		Context("with method type not matching interface method type", func() {
			BeforeEach(func() {
				method = usage.NewInterface3Mock().AnotherMethod
			})

			It("should panic", func() {
				act()
				Expect(act).To(PanicWith(BeAssignableToTypeOf(detox.InterfaceMethodMismatchError{})))
			})
		})
	})
})
