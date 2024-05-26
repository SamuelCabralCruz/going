//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	"github.com/SamuelCabralCruz/went/fn"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var cut fixture.Interface1Mock

	BeforeEach(func() {
		cut = fixture.NewInterface1Mock()
	})

	AfterEach(func() {
		cut.Reset()
	})

	DescribeFunction(cut.Default, func() {
		Context("without default implementation", func() {
			DescribeTable("should panic on invocation", func(act func()) {
				Expect(act).To(PanicWith(BeAssignableToTypeOf(fake.MissingImplementationError{})))
			},
				CreateTableEntries([]string{"act"},
					[]any{func() { cut.NoArgNoReturn() }},
					[]any{func() { _ = cut.NoArgSingleReturn() }},
					[]any{func() { _, _, _ = cut.NoArgMultipleReturns() }},
					[]any{func() { cut.SingleArgNoReturn("") }},
					[]any{func() { _ = cut.SingleArgSingleReturn("") }},
					[]any{func() { _, _ = cut.SingleArgMultipleReturns("") }},
					[]any{func() { cut.MultipleArgsNoReturn(1, true, []byte("arg3")) }},
					[]any{func() { _ = cut.MultipleArgsSingleReturn(1, true) }},
					[]any{func() { _, _ = cut.MultipleArgsMultipleReturns(0.2, uint8(2)) }},
				),
			)
		})

		Context("with default implementation", func() {
			BeforeEach(func() {
				cut.Default(fixture.Implementation1{})
			})

			DescribeTable("should invoke default implementation", func(act func()) {
				Expect(act).NotTo(Panic())
			},
				CreateTableEntries([]string{"act"},
					[]any{func() { cut.NoArgNoReturn() }},
					[]any{func() { _ = cut.NoArgSingleReturn() }},
					[]any{func() { _, _, _ = cut.NoArgMultipleReturns() }},
					[]any{func() { cut.SingleArgNoReturn("") }},
					[]any{func() { _ = cut.SingleArgSingleReturn("") }},
					[]any{func() { _, _ = cut.SingleArgMultipleReturns("") }},
					[]any{func() { cut.MultipleArgsNoReturn(1, true, []byte("arg3")) }},
					[]any{func() { _ = cut.MultipleArgsSingleReturn(1, true) }},
					[]any{func() { _, _ = cut.MultipleArgsMultipleReturns(0.2, uint8(2)) }},
				),
			)
		})

		Context("with already existing fake", func() {
			BeforeEach(func() {
				_ = fn.Prevent(cut.NoArgNoReturn)
				cut.Default(fixture.Implementation1{})
			})

			It("should propagate newly registered default implementation", func() {
				Expect(cut.NoArgNoReturn).NotTo(Panic())
			})
		})

		Context("with registered method", func() {
			cut := fixture.NewInterface3Mock()
			var registeredCalled bool
			var defaultImpl *fixture.Implementation3

			BeforeEach(func() {
				detox.When(cut.Detox, cut.Method).Call(func(_ string) {
					registeredCalled = true
				})
				defaultImpl = &fixture.Implementation3{}
				cut.Default(defaultImpl)
			})

			It("should have priority over default implementation", func() {
				cut.Method("input value")

				Expect(defaultImpl.Called).To(BeFalse())
				Expect(registeredCalled).To(BeTrue())
			})
		})

		Context("with already registered default implementation", func() {
			cut := fixture.NewInterface3Mock()
			var defaultImpl1 *fixture.Implementation3
			var defaultImpl2 *fixture.Implementation3

			BeforeEach(func() {
				defaultImpl1 = &fixture.Implementation3{}
				defaultImpl2 = &fixture.Implementation3{}
				cut.Default(defaultImpl1)
				cut.Default(defaultImpl2)
			})

			It("should override previously registered default implementation", func() {
				cut.Method("input value")

				Expect(defaultImpl1.Called).To(BeFalse())
				Expect(defaultImpl2.Called).To(BeTrue())
			})
		})
	})
})
