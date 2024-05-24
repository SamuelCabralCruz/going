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
	var cut usage.Interface3Mock
	var mocked detox.Mocked[func(string)]
	var defaultImplementation *usage.Implementation3
	var persistentCalled bool
	var ephemeralCalled bool
	var persistentConditionalCalled bool
	var ephemeralConditionalCalled bool
	var inputArg string

	act := func() {
		cut.Method(inputArg)
	}

	BeforeEach(func() {
		cut = usage.NewInterface3Mock()
		mocked = detox.When(cut.Detox, cut.Method)
		persistentCalled = false
		ephemeralCalled = false
		persistentConditionalCalled = false
		ephemeralConditionalCalled = false
		inputArg = "some input value"
	})

	AfterEach(func() {
		cut.Reset()
	})

	givenDefaultImplementation := func() {
		defaultImplementation = &usage.Implementation3{}
		cut.Default(defaultImplementation)
	}

	givenPersistentRegistration := func() {
		mocked.Call(func(_ string) {
			persistentCalled = true
		})
	}

	givenEphemeralRegistration := func() {
		mocked.CallOnce(func(_ string) {
			ephemeralCalled = true
		})
	}

	givenPersistentConditionalRegistration := func() {
		inputArg = "some conditional value"
		mocked.WithArgs(inputArg).Call(func(_ string) {
			persistentConditionalCalled = true
		})
	}

	givenEphemeralConditionalRegistration := func() {
		inputArg = "some conditional value"
		mocked.WithArgs(inputArg).CallOnce(func(_ string) {
			ephemeralConditionalCalled = true
		})
	}

	assertInvocation := func(p bool, e bool, pc bool, ec bool) {
		act()

		Expect(defaultImplementation.Called).To(BeFalse())
		Expect(persistentCalled).To(Equal(p))
		Expect(ephemeralCalled).To(Equal(e))
		Expect(persistentConditionalCalled).To(Equal(pc))
		Expect(ephemeralConditionalCalled).To(Equal(ec))
	}

	DescribeFunction(detox.Mocked[any].ResolveForArgs, func() {
		Context("with default implementation", func() {
			BeforeEach(givenDefaultImplementation)

			Context("with persistent registration", func() {
				BeforeEach(givenPersistentRegistration)

				It("should call persistent registration", func() {
					assertInvocation(true, false, false, false)
				})
			})

			Context("with ephemeral registration", func() {
				BeforeEach(givenEphemeralRegistration)

				It("should call ephemeral registration", func() {
					assertInvocation(false, true, false, false)
				})
			})

			Context("with persistent conditional registration", func() {
				BeforeEach(givenPersistentConditionalRegistration)

				It("should call persistent conditional registration", func() {
					assertInvocation(false, false, true, false)
				})
			})

			Context("with ephemeral conditional registration", func() {
				BeforeEach(givenEphemeralConditionalRegistration)

				It("should call ephemeral conditional registration", func() {
					assertInvocation(false, false, false, true)
				})
			})
		})

		Context("with persistent registration", func() {
			BeforeEach(givenPersistentRegistration)

			Context("with ephemeral registration", func() {
				BeforeEach(givenEphemeralRegistration)

				It("should call ephemeral registration", func() {
					assertInvocation(false, true, false, false)
				})
			})

			Context("with persistent conditional registration", func() {
				BeforeEach(givenPersistentConditionalRegistration)

				It("should call persistent conditional registration", func() {
					assertInvocation(false, false, true, false)
				})
			})

			Context("with ephemeral conditional registration", func() {
				BeforeEach(givenEphemeralConditionalRegistration)

				It("should call ephemeral conditional registration", func() {
					assertInvocation(false, false, false, true)
				})
			})
		})

		Context("with ephemeral registration", func() {
			BeforeEach(givenEphemeralRegistration)

			Context("with persistent conditional registration", func() {
				BeforeEach(givenPersistentConditionalRegistration)

				It("should call persistent conditional registration", func() {
					assertInvocation(false, false, true, false)
				})
			})

			Context("with ephemeral conditional registration", func() {
				BeforeEach(givenEphemeralConditionalRegistration)

				It("should call ephemeral conditional registration", func() {
					assertInvocation(false, false, false, true)
				})
			})
		})

		Context("with persistent conditional registration", func() {
			BeforeEach(givenPersistentConditionalRegistration)

			Context("with ephemeral conditional registration", func() {
				BeforeEach(givenEphemeralConditionalRegistration)

				It("should call ephemeral conditional registration", func() {
					assertInvocation(false, false, false, true)
				})
			})
		})
	})
})
