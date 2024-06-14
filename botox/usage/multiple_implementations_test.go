//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/container"
	"github.com/SamuelCabralCruz/went/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.MustResolveAll[any], func() {
	var impl1 fixture.SomeImplementation1
	var impl2 fixture.SomeImplementation2
	var impl3 fixture.SomeImplementation3
	var observed []fixture.SomeInterface

	act := func() {
		observed = botox.MustResolveAll[fixture.SomeInterface]()
	}

	BeforeEach(func() {
		impl1 = fixture.SomeImplementation1{}
		impl2 = fixture.SomeImplementation2{}
		impl3 = fixture.SomeImplementation3{}
	})

	AfterEach(func() {
		botox.Reset()
	})

	Context("with implementations registered on their own type", func() {
		BeforeEach(func() {
			botox.RegisterInstance[fixture.SomeImplementation1](impl1)
			botox.RegisterInstance[fixture.SomeImplementation2](impl2)
			botox.RegisterInstance[fixture.SomeImplementation3](impl3)
		})

		It("should not resolve any implementation", func() {
			Expect(act).To(PanicWith(BeAssignableToTypeOf(container.NoCandidateFoundError{})))
		})
	})

	Context("with implementations registered on the interface", func() {
		BeforeEach(func() {
			botox.RegisterInstance[fixture.SomeInterface](impl1)
			botox.RegisterInstance[fixture.SomeInterface](impl2)
			botox.RegisterInstance[fixture.SomeInterface](impl3)
		})

		It("should resolve all implementations", func() {
			act()

			Expect(observed).To(HaveLen(3))
			Expect(observed).To(Equal([]fixture.SomeInterface{impl1, impl2, impl3}))
		})
	})
})
