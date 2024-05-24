//go:build test

package usage_test

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/usage"
	. "github.com/SamuelCabralCruz/went/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var cut usage.Interface1Mock
	var mocked detox.Mocked[func(string) string]

	BeforeEach(func() {
		cut = usage.NewInterface1Mock()
		mocked = detox.When(cut.Detox, cut.SingleArgSingleReturn)
	})

	AfterEach(func() {
		cut.Reset()
	})

	DescribeFunction(detox.Mocked[any].WithArgs, func() {
		DescribeFunction(detox.Mocked[any].Call, func() {
			var mockedArg string
			var providedArg string
			var observed string

			act := func() {
				observed = cut.SingleArgSingleReturn(providedArg)
			}

			BeforeEach(func() {
				mockedArg = "some input value"
				mocked.WithArgs(mockedArg).Call(func(s string) string {
					return fmt.Sprintf("this return has been mocked - %s", s)
				})
			})

			Context("with matching arguments invocation", func() {
				BeforeEach(func() {
					providedArg = mockedArg
				})

				It("should register fake implementation", func() {
					act()

					Expect(observed).To(Equal("this return has been mocked - some input value"))
				})

				It("should be persistent", func() {
					act()

					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).NotTo(Panic())
					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).NotTo(Panic())
					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).NotTo(Panic())
					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).NotTo(Panic())
					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).NotTo(Panic())
				})
			})

			Context("with non matching arguments invocation", func() {
				BeforeEach(func() {
					providedArg = "non mocked value"
				})

				It("should not resolve fake implementation", func() {
					Expect(act).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				})
			})
		})
	})
})
