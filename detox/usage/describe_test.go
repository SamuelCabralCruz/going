//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/going/detox"
	"github.com/SamuelCabralCruz/going/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var cut fixture.Interface1Mock
	var mocked1 detox.Mocked[func(string) string]
	var mocked2 detox.Mocked[func()]
	var mocked3 detox.Mocked[func(string)]

	BeforeEach(func() {
		cut = fixture.NewInterface1Mock()

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
			Expect(observed1).To(MatchRegexp("^Interface1\\.SingleArgSingleReturn \\(.*/going/detox/usage/describe_test\\.go \\[20\\]\\)$"))
			Expect(observed2).To(MatchRegexp("^Interface1\\.NoArgNoReturn \\(.*/going/detox/usage/describe_test\\.go \\[20\\]\\)$"))
			Expect(observed3).To(MatchRegexp("^Interface1\\.SingleArgNoReturn \\(.*/going/detox/usage/describe_test\\.go \\[20\\]\\)$"))
		})
	})
})
