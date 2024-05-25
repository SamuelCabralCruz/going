//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/usage"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var cut usage.Interface1Mock
	var mocked1 detox.Mocked[func(string) string]
	var mocked2 detox.Mocked[func()]
	var mocked3 detox.Mocked[func(string)]

	BeforeEach(func() {
		cut = usage.NewInterface1Mock()

		mocked1 = detox.When(cut.Detox, cut.SingleArgSingleReturn)
		mocked2 = detox.When(cut.Detox, cut.NoArgNoReturn)
		mocked3 = detox.When(cut.Detox, cut.SingleArgNoReturn)

		mocked1.Call(func(_ string) string { return "fake1" })
		mocked2.Call(func() {})
		mocked1.WithArgs("fake3").Call(func(_ string) string { return "fake3" })
		mocked3.Call(func(_ string) {})
		mocked3.WithArgs("fake4").Call(func(_ string) {})
	})

	AfterEach(func() {
		cut.Reset()
	})

	DescribeFunction(cut.Reset, func() {
		act := func() {
			cut.Reset()
		}

		Context("with registered fakes", func() {
			It("should clear all implementations", func() {
				act()

				Expect(func() { _ = cut.SingleArgSingleReturn("any arg") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.NoArgNoReturn() }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { _ = cut.SingleArgSingleReturn("fake3") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.SingleArgNoReturn("any arg") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.SingleArgNoReturn("fake4") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
			})
		})

		Context("with registered spies", func() {
			BeforeEach(func() {
				cut.SingleArgSingleReturn("1st")
				cut.SingleArgSingleReturn("2nd")
				cut.NoArgNoReturn()
				cut.NoArgNoReturn()
				cut.NoArgNoReturn()
				cut.SingleArgSingleReturn("6th")
				cut.SingleArgSingleReturn("7th")
				cut.SingleArgNoReturn("8th")
				cut.SingleArgNoReturn("9th")
				cut.SingleArgNoReturn("10th")
			})

			It("should clear all invocations", func() {
				act()

				Expect(mocked1.Assert().HasBeenCalled()).To(BeFalse())
				Expect(mocked2.Assert().HasBeenCalled()).To(BeFalse())
				Expect(mocked3.Assert().HasBeenCalled()).To(BeFalse())
			})
		})

		Context("with default implementation", func() {
			BeforeEach(func() {
				cut.Default(usage.Implementation1{})
			})

			It("should clear default implementation", func() {
				act()

				Expect(func() { _ = cut.SingleArgSingleReturn("any arg") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.NoArgNoReturn() }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { _ = cut.SingleArgSingleReturn("fake3") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.SingleArgNoReturn("any arg") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { cut.SingleArgNoReturn("fake4") }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
			})
		})
	})

	DescribeFunction(detox.Mocked[any].Reset, func() {
		act := func() {
			mocked2.Reset()
		}

		Context("with registered fakes", func() {
			It("should only clear targeted fake", func() {
				act()

				Expect(func() { _ = cut.SingleArgSingleReturn("any arg") }).NotTo(Panic())
				Expect(func() { cut.NoArgNoReturn() }).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
				Expect(func() { _ = cut.SingleArgSingleReturn("fake3") }).NotTo(Panic())
				Expect(func() { cut.SingleArgNoReturn("any arg") }).NotTo(Panic())
				Expect(func() { cut.SingleArgNoReturn("fake4") }).NotTo(Panic())
			})
		})

		Context("with registered spies", func() {
			BeforeEach(func() {
				cut.SingleArgSingleReturn("1st")
				cut.SingleArgSingleReturn("2nd")
				cut.NoArgNoReturn()
				cut.NoArgNoReturn()
				cut.NoArgNoReturn()
				cut.SingleArgSingleReturn("6th")
				cut.SingleArgSingleReturn("7th")
				cut.SingleArgNoReturn("8th")
				cut.SingleArgNoReturn("9th")
				cut.SingleArgNoReturn("10th")
			})

			It("should only clear targeted spy", func() {
				act()

				Expect(mocked1.Assert().HasBeenCalled()).To(BeTrue())
				Expect(mocked2.Assert().HasBeenCalled()).To(BeFalse())
				Expect(mocked3.Assert().HasBeenCalled()).To(BeTrue())
			})
		})
	})
})
