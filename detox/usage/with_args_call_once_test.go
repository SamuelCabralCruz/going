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
		DescribeFunction(detox.Mocked[any].CallOnce, func() {
			var mockedArg string
			var providedArg string
			var observed string

			act := func() {
				mocked.WithArgs(mockedArg).CallOnce(func(s string) string {
					return fmt.Sprintf("this return has been mocked - %s", s)
				})
				observed = cut.SingleArgSingleReturn(providedArg)
			}

			BeforeEach(func() {
				mockedArg = "some input value"
				providedArg = mockedArg
			})

			Context("with matching arguments invocation", func() {
				It("should register fake implementation", func() {
					act()

					Expect(observed).To(Equal("this return has been mocked - some input value"))
				})

				It("should be ephemeral", func() {
					act()

					Expect(func() { cut.SingleArgSingleReturn(mockedArg) }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
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

			Context("with already registered ephemeral conditional implementation", func() {
				BeforeEach(func() {
					mocked.WithArgs(mockedArg).CallOnce(func(_ string) string {
						return "already registered"
					})
				})

				It("should accumulate ephemeral implementations and resolve them in order", func() {
					act()

					Expect(observed).To(Equal("already registered"))
					Expect(cut.SingleArgSingleReturn(providedArg)).To(Equal("this return has been mocked - some input value"))
					Expect(func() { cut.SingleArgSingleReturn(providedArg) }).To(Panic())
				})
			})
		})
	})
})
