//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	var nestedMock fixture.Interface1Mock
	var parentMock fixture.Interface2Mock
	var mockedProvider detox.Mocked[func() fixture.Interface1]
	var mockedNested detox.Mocked[func(string) string]
	var observed string
	var expected string

	act := func() {
		observed = parentMock.ReturnAnotherInterface().SingleArgSingleReturn("some input value")
	}

	BeforeEach(func() {
		nestedMock = fixture.NewInterface1Mock()
		parentMock = fixture.NewInterface2Mock()
		mockedProvider = detox.When(parentMock.Detox, parentMock.ReturnAnotherInterface)
		mockedNested = detox.When(nestedMock.Detox, nestedMock.SingleArgSingleReturn)
		mockedProvider.Call(func() fixture.Interface1 {
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
