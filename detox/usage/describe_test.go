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
	var cut usage.Interface1Mock
	var mocked1 detox.Mocked[func(string) string]
	var mocked2 detox.Mocked[func()]
	var mocked3 detox.Mocked[func(string)]

	BeforeEach(func() {
		cut = usage.NewInterface1Mock()

		mocked1 = detox.When(cut.Detox, cut.SingleArgSingleReturn)
		mocked2 = detox.When(cut.Detox, cut.NoArgNoReturn)
		mocked3 = detox.When(cut.Detox, cut.SingleArgNoReturn)
	})

	AfterEach(func() {
		cut.Reset()
	})

	DescribeFunction(detox.Mocked[any].Describe, func() {
		It("should be specific to method level", func() {
			observed1 := mocked1.Describe()
			observed2 := mocked2.Describe()
			observed3 := mocked3.Describe()

			Expect(observed1).NotTo(Equal(observed2))
			Expect(observed1).NotTo(Equal(observed3))
			Expect(observed1).To(MatchRegexp("^Interface1\\.SingleArgSingleReturn \\(.*/went/detox/usage/describe_test\\.go \\[20\\]\\)$"))
			Expect(observed2).To(MatchRegexp("^Interface1\\.NoArgNoReturn \\(.*/went/detox/usage/describe_test\\.go \\[20\\]\\)$"))
			Expect(observed3).To(MatchRegexp("^Interface1\\.SingleArgNoReturn \\(.*/went/detox/usage/describe_test\\.go \\[20\\]\\)$"))
		})
	})
})
