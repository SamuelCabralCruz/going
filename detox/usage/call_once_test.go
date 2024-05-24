//go:build test

package usage_test

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/detox"
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

	DescribeFunction(detox.Mocked[any].CallOnce, func() {
		var observed string

		act := func() {
			mocked.CallOnce(func(s string) string {
				return fmt.Sprintf("this return has been mocked - %s", s)
			})
			observed = cut.SingleArgSingleReturn("some input value")
		}

		It("should resolve fake implementation", func() {
			act()

			Expect(observed).To(Equal("this return has been mocked - some input value"))
		})

		It("should be ephemeral", func() {
			act()

			Expect(func() { cut.SingleArgSingleReturn("some input value") }).To(Panic())
		})

		Context("with already registered ephemeral implementation", func() {
			BeforeEach(func() {
				mocked.CallOnce(func(_ string) string {
					return "already registered"
				})
			})

			It("should accumulate ephemeral implementations and resolve them in order", func() {
				act()

				Expect(observed).To(Equal("already registered"))
				Expect(cut.SingleArgSingleReturn("some other input value")).To(Equal("this return has been mocked - some other input value"))
				Expect(func() { cut.SingleArgSingleReturn("some input") }).To(Panic())
			})
		})
	})
})
