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
	var cut fixture.Stateless

	BeforeEach(func() {
		cut = fixture.NewStateless()
	})

	AfterEach(func() {
		botox.Reset()
	})

	assert := func(act func(), shouldNotPanic bool) {
		if shouldNotPanic {
			Expect(act).NotTo(Panic())
		} else {
			Expect(act).To(Panic())
		}
	}

	assertAll := func(shouldResolveValue bool, shouldResolveReference bool, shouldResolveReferenceOnReference bool) {
		assert(func() { botox.MustResolve[fixture.Stateless]() }, shouldResolveValue)
		assert(func() { botox.MustResolve[*fixture.Stateless]() }, shouldResolveReference)
		assert(func() { botox.MustResolve[**fixture.Stateless]() }, shouldResolveReferenceOnReference)
	}

	Context("with value instance registration", func() {
		BeforeEach(func() {
			botox.RegisterInstance(cut)
		})

		It("should resolve value", func() {
			assertAll(true, false, false)
		})
	})

	Context("with reference instance registration", func() {
		BeforeEach(func() {
			botox.RegisterInstance(&cut)
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with value singleton instance registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonInstance(cut)
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with reference singleton instance registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonInstance(&cut)
		})

		It("should resolve reference on reference", func() {
			assertAll(false, false, true)
		})
	})

	Context("with value supplier registration", func() {
		BeforeEach(func() {
			botox.RegisterSupplier(func() fixture.Stateless {
				return cut
			})
		})

		It("should resolve value", func() {
			assertAll(true, false, false)
		})
	})

	Context("with reference supplier registration", func() {
		BeforeEach(func() {
			botox.RegisterSupplier(func() *fixture.Stateless {
				return &cut
			})
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with value singleton supplier registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonSupplier(func() fixture.Stateless {
				return cut
			})
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with reference singleton supplier registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonSupplier(func() *fixture.Stateless {
				return &cut
			})
		})

		It("should resolve reference on reference", func() {
			assertAll(false, false, true)
		})
	})

	Context("with value producer registration", func() {
		BeforeEach(func() {
			botox.RegisterProducer(func() (fixture.Stateless, error) {
				return cut, nil
			})
		})

		It("should resolve value", func() {
			assertAll(true, false, false)
		})
	})

	Context("with reference producer registration", func() {
		BeforeEach(func() {
			botox.RegisterProducer(func() (*fixture.Stateless, error) {
				return &cut, nil
			})
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with value singleton producer registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonProducer(func() (fixture.Stateless, error) {
				return cut, nil
			})
		})

		It("should resolve reference", func() {
			assertAll(false, true, false)
		})
	})

	Context("with reference singleton producer registration", func() {
		BeforeEach(func() {
			botox.RegisterSingletonProducer(func() (*fixture.Stateless, error) {
				return &cut, nil
			})
		})

		It("should resolve reference on reference", func() {
			assertAll(false, false, true)
		})
	})
})
