//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/usage"
	. "github.com/SamuelCabralCruz/went/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var nestedMock usage.Interface1Mock
	var parentMock usage.Interface2Mock
	var mockedProvider detox.Mocked[func() usage.Interface1]
	var mockedNested detox.Mocked[func(string) string]
	var observed string
	var expected string

	act := func() {
		observed = parentMock.ReturnAnotherInterface().SingleArgSingleReturn("some input value")
	}

	BeforeEach(func() {
		nestedMock = usage.NewInterface1Mock()
		parentMock = usage.NewInterface2Mock()
		mockedProvider = detox.When(parentMock.Detox, parentMock.ReturnAnotherInterface)
		mockedNested = detox.When(nestedMock.Detox, nestedMock.SingleArgSingleReturn)
		mockedProvider.Call(func() usage.Interface1 {
			return nestedMock
		})
		expected = "property mocked"
		mockedNested.CallOnce(func(_ string) string {
			return expected
		})
	})

	AfterEach(func() {
		nestedMock.Reset()
		parentMock.Reset()
	})

	It("should be able to chain mocks", func() {
		act()

		Expect(observed).To(Equal(expected))
	})
})
