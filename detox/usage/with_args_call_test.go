//go:build test

package usage_test

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/usage"
	. "github.com/SamuelCabralCruz/went/kinggo"
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
				mocked.WithArgs(mockedArg).Call(func(s string) string {
					return fmt.Sprintf("this return has been mocked - %s", s)
				})
				observed = cut.SingleArgSingleReturn(providedArg)
			}

			BeforeEach(func() {
				mockedArg = "some input value"
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

			Context("with already registered persistent conditional implementation", func() {
				BeforeEach(func() {
					providedArg = mockedArg
				})

				Context("with same condition", func() {
					BeforeEach(func() {
						mocked.WithArgs(mockedArg).Call(func(s string) string {
							return "already registered"
						})
					})

					It("should override", func() {
						act()

						Expect(observed).To(Equal("this return has been mocked - some input value"))
					})
				})

				Context("with different condition", func() {
					var otherMockedArg string

					BeforeEach(func() {
						otherMockedArg = "some other arg"
						mocked.WithArgs(otherMockedArg).Call(func(s string) string {
							return "already registered"
						})
					})

					It("should not override", func() {
						act()

						Expect(observed).To(Equal("this return has been mocked - some input value"))
						Expect(cut.SingleArgSingleReturn(otherMockedArg)).To(Equal("already registered"))
					})
				})
			})
		})
	})
})
