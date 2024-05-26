//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.MustResolve[any], func() {
	var cut fixture.Stateful

	BeforeEach(func() {
		cut = fixture.Stateful{}
	})

	AfterEach(func() {
		botox.Clear()
	})

	Context("with non singleton value registration", func() {
		BeforeEach(func() {
			botox.RegisterInstance(cut)
		})

		Context("without prior mutation", func() {
			var resolved1 fixture.Stateful
			var resolved2 fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[fixture.Stateful]()
				resolved2 = botox.MustResolve[fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			It("should isolate resolved instances", func() {
				act()

				Expect(resolved1.Count()).To(Equal(2))
				Expect(resolved2.Count()).To(Equal(1))
				Expect(cut.Count()).To(Equal(0))
			})
		})

		Context("with prior mutations", func() {
			var resolved1 fixture.Stateful
			var resolved2 fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[fixture.Stateful]()
				resolved2 = botox.MustResolve[fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			BeforeEach(func() {
				resolved := botox.MustResolve[fixture.Stateful]()
				resolved.Mutate()
				resolved.Mutate()
			})

			It("should resolve fresh instances", func() {
				act()

				Expect(resolved1.Count()).To(Equal(2))
				Expect(resolved2.Count()).To(Equal(1))
				Expect(cut.Count()).To(Equal(0))
			})
		})
	})

	Context("with singleton value registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonInstance(cut)
		})

		Context("without prior mutation", func() {
			var resolved1 *fixture.Stateful
			var resolved2 *fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[*fixture.Stateful]()
				resolved2 = botox.MustResolve[*fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			It("should share state between all instances", func() {
				act()

				Expect(resolved1.Count()).To(Equal(3))
				Expect(resolved2.Count()).To(Equal(3))
				Expect(cut.Count()).To(Equal(0))
			})
		})

		Context("with prior mutations", func() {
			var resolved1 *fixture.Stateful
			var resolved2 *fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[*fixture.Stateful]()
				resolved2 = botox.MustResolve[*fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			BeforeEach(func() {
				resolved := botox.MustResolve[*fixture.Stateful]()
				resolved.Mutate()
				resolved.Mutate()
			})

			It("should resolve instances from previous state", func() {
				act()

				Expect(resolved1.Count()).To(Equal(5))
				Expect(resolved2.Count()).To(Equal(5))
				Expect(cut.Count()).To(Equal(0))
			})
		})
	})

	Context("with non singleton reference registration", func() {
		BeforeEach(func() {
			botox.RegisterInstance(&cut)
		})

		Context("without prior mutation", func() {
			var resolved1 *fixture.Stateful
			var resolved2 *fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[*fixture.Stateful]()
				resolved2 = botox.MustResolve[*fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			It("should share state between all instances", func() {
				act()

				Expect(resolved1.Count()).To(Equal(3))
				Expect(resolved2.Count()).To(Equal(3))
				Expect(cut.Count()).To(Equal(3))
			})
		})

		Context("with prior mutations", func() {
			var resolved1 *fixture.Stateful
			var resolved2 *fixture.Stateful

			act := func() {
				resolved1 = botox.MustResolve[*fixture.Stateful]()
				resolved2 = botox.MustResolve[*fixture.Stateful]()
				resolved1.Mutate()
				resolved2.Mutate()
				resolved1.Mutate()
			}

			BeforeEach(func() {
				resolved := botox.MustResolve[*fixture.Stateful]()
				resolved.Mutate()
				resolved.Mutate()
			})

			It("should resolve instances from previous state", func() {
				act()

				Expect(resolved1.Count()).To(Equal(5))
				Expect(resolved2.Count()).To(Equal(5))
				Expect(cut.Count()).To(Equal(5))
			})
		})
	})
})
