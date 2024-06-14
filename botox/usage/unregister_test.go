//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.Unregister[any], func() {
	var impl1 fixture.SomeImplementation1
	var impl2 fixture.SomeImplementation2
	var impl3 fixture.SomeImplementation3
	var impl4 fixture.Stateless

	BeforeEach(func() {
		impl1 = fixture.SomeImplementation1{}
		impl2 = fixture.SomeImplementation2{}
		impl3 = fixture.SomeImplementation3{}
		impl4 = fixture.Stateless{}
		botox.RegisterInstance[fixture.SomeInterface](impl1)
		botox.RegisterInstance[fixture.SomeInterface](impl2)
		botox.RegisterInstance[fixture.SomeInterface](impl3)
		botox.RegisterInstance(impl4)
	})

	AfterEach(func() {
		botox.Reset()
	})

	It("should unregister all tokens of the specified type", func() {
		observed := botox.MustResolveAll[fixture.SomeInterface]()
		Expect(observed).To(HaveLen(3))

		botox.Unregister[fixture.SomeImplementation1]()

		observed = botox.MustResolveAll[fixture.SomeInterface]()
		Expect(observed).To(HaveLen(3))

		botox.Unregister[fixture.SomeInterface]()

		observed, _ = botox.ResolveAll[fixture.SomeInterface]()
		Expect(observed).To(HaveLen(0))

		Expect(func() { _ = botox.MustResolve[fixture.Stateless] }).NotTo(Panic())
	})
})
