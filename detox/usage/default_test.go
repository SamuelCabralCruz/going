//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/internal/fake"
	"github.com/SamuelCabralCruz/went/detox/usage"
	"github.com/SamuelCabralCruz/went/fn"
	. "github.com/SamuelCabralCruz/went/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var cut usage.Interface1Mock

	BeforeEach(func() {
		cut = usage.NewInterface1Mock()
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
					[]any{func() { cut.MultipleArgsNoReturn(1, true, []byte("arg3")) }},
					[]any{func() { _, _ = cut.MultipleArgsMultipleReturns(0.2, uint8(2)) }},
				),
			)
		})

		Context("with default implementation", func() {
			BeforeEach(func() {
				cut.Default(usage.Implementation1{})
			})

			DescribeTable("should invoke default implementation", func(act func()) {
				Expect(act).NotTo(Panic())
			},
				CreateTableEntries([]string{"act"},
					[]any{func() { cut.NoArgNoReturn() }},
					[]any{func() { _ = cut.NoArgSingleReturn() }},
					[]any{func() { _, _, _ = cut.NoArgMultipleReturns() }},
					[]any{func() { cut.SingleArgNoReturn("") }},
					[]any{func() { cut.MultipleArgsNoReturn(1, true, []byte("arg3")) }},
					[]any{func() { _, _ = cut.MultipleArgsMultipleReturns(0.2, uint8(2)) }},
				),
			)
		})

		Context("with already existing fake", func() {
			BeforeEach(func() {
				_ = fn.Prevent(cut.NoArgNoReturn)
				cut.Default(usage.Implementation1{})
			})

			It("should propagate newly registered default implementation", func() {
				Expect(cut.NoArgNoReturn).NotTo(Panic())
			})
		})

		Context("with registered method", func() {
			cut := usage.NewInterface3Mock()
			var registeredCalled bool
			var defaultImpl *usage.Implementation3

			BeforeEach(func() {
				detox.When(cut.Detox, cut.Method).Call(func() {
					registeredCalled = true
				})
				defaultImpl = &usage.Implementation3{}
				cut.Default(defaultImpl)
			})

			It("should have priority over default implementation", func() {
				cut.Method()

				Expect(defaultImpl.Called).To(BeFalse())
				Expect(registeredCalled).To(BeTrue())
			})
		})

		Context("with already registered default implementation", func() {
			cut := usage.NewInterface3Mock()
			var defaultImpl1 *usage.Implementation3
			var defaultImpl2 *usage.Implementation3

			BeforeEach(func() {
				defaultImpl1 = &usage.Implementation3{}
				defaultImpl2 = &usage.Implementation3{}
				cut.Default(defaultImpl1)
				cut.Default(defaultImpl2)
			})

			It("should override previously registered default implementation", func() {
				cut.Method()

				Expect(defaultImpl1.Called).To(BeFalse())
				Expect(defaultImpl2.Called).To(BeTrue())
			})
		})
	})
})
