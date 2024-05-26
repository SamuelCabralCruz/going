//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.RegisterSingletonInstance[any], func() {
	var cut fixture.Stateless

	register := func() {
		botox.RegisterSingletonInstance(cut)
	}

	registerMultiple := func(instances ...fixture.Stateless) {
		for _, instance := range instances {
			botox.RegisterSingletonInstance(instance)
		}
	}

	BeforeEach(func() {
		cut = fixture.NewStateless()
	})

	AfterEach(func() {
		botox.Clear()
	})

	DescribeFunction(botox.MustResolve[any], func() {
		var observed *fixture.Stateless

		resolve := func() {
			observed = botox.MustResolve[*fixture.Stateless]()
		}

		Context("without registration", func() {
			It("should panic", func() {
				Expect(resolve).To(PanicWith(BeAssignableToTypeOf(botox.NoCandidateFoundError{})))
			})
		})

		Context("with registration", func() {
			BeforeEach(func() {
				register()
			})

			It("should return registered instance", func() {
				resolve()

				Expect(observed).To(Equal(&cut))
			})
		})

		Context("with multiple registrations", func() {
			BeforeEach(func() {
				registerMultiple(cut, fixture.NewStateless())
			})

			It("should panic", func() {
				Expect(resolve).To(PanicWith(BeAssignableToTypeOf(botox.TooManyCandidatesFoundError{})))
			})
		})
	})

	DescribeFunction(botox.Resolve[any], func() {
		var observedInstance *fixture.Stateless
		var observedError error

		resolve := func() {
			observedInstance, observedError = botox.Resolve[*fixture.Stateless]()
		}

		Context("without registration", func() {
			It("should return error", func() {
				resolve()

				Expect(observedInstance).To(BeZero())
				Expect(observedError).To(BeAssignableToTypeOf(botox.NoCandidateFoundError{}))
			})
		})

		Context("with registration", func() {
			BeforeEach(func() {
				register()
			})

			It("should return registered instance", func() {
				resolve()

				Expect(observedInstance).To(Equal(&cut))
				Expect(observedError).To(BeNil())
			})
		})

		Context("with multiple registrations", func() {
			otherRegistration1 := fixture.NewStateless()
			otherRegistration2 := fixture.NewStateless()

			BeforeEach(func() {
				registerMultiple(cut, otherRegistration1, otherRegistration2)
			})

			It("should return error", func() {
				resolve()

				Expect(observedInstance).To(BeZero())
				Expect(observedError).To(BeAssignableToTypeOf(botox.TooManyCandidatesFoundError{}))
			})
		})
	})

	DescribeFunction(botox.MustResolveAll[any], func() {
		var observed []*fixture.Stateless

		resolve := func() {
			observed = botox.MustResolveAll[*fixture.Stateless]()
		}

		Context("without registration", func() {
			It("should panic", func() {
				Expect(resolve).To(PanicWith(BeAssignableToTypeOf(botox.NoCandidateFoundError{})))
			})
		})

		Context("with registration", func() {
			BeforeEach(func() {
				register()
			})

			It("should return registered instance", func() {
				resolve()

				Expect(observed).To(HaveLen(1))
				Expect(observed).To(Equal([]*fixture.Stateless{&cut}))
			})
		})

		Context("with multiple registrations", func() {
			var otherRegistration1 fixture.Stateless
			var otherRegistration2 fixture.Stateless

			BeforeEach(func() {
				otherRegistration1 = fixture.NewStateless()
				otherRegistration2 = fixture.NewStateless()
				registerMultiple(cut, otherRegistration1, otherRegistration2)
			})

			It("should return all registered instances", func() {
				resolve()

				Expect(observed).To(HaveLen(3))
				Expect(observed).To(Equal([]*fixture.Stateless{&cut, &otherRegistration1, &otherRegistration2}))
			})
		})
	})

	DescribeFunction(botox.ResolveAll[any], func() {
		var observedInstances []*fixture.Stateless
		var observedError error

		resolve := func() {
			observedInstances, observedError = botox.ResolveAll[*fixture.Stateless]()
		}

		Context("without registration", func() {
			It("should return error", func() {
				resolve()

				Expect(observedInstances).To(HaveLen(0))
				Expect(observedError).To(BeAssignableToTypeOf(botox.NoCandidateFoundError{}))
			})
		})

		Context("with registration", func() {
			BeforeEach(func() {
				register()
			})

			It("should return registered instance", func() {
				resolve()

				Expect(observedInstances).To(HaveLen(1))
				Expect(observedInstances).To(Equal([]*fixture.Stateless{&cut}))
				Expect(observedError).To(BeNil())
			})
		})

		Context("with multiple registrations", func() {
			var otherRegistration1 fixture.Stateless
			var otherRegistration2 fixture.Stateless

			BeforeEach(func() {
				otherRegistration1 = fixture.NewStateless()
				otherRegistration2 = fixture.NewStateless()
				registerMultiple(cut, otherRegistration1, otherRegistration2)
			})

			It("should return all registered instances", func() {
				resolve()

				Expect(observedInstances).To(HaveLen(3))
				Expect(observedInstances).To(Equal([]*fixture.Stateless{&cut, &otherRegistration1, &otherRegistration2}))
				Expect(observedError).To(BeNil())
			})
		})
	})
})
