//go:build test

package usage_test

import (
	"github.com/SamuelCabralCruz/going/botox"
	"github.com/SamuelCabralCruz/going/botox/internal/it"
	"github.com/SamuelCabralCruz/going/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(it.RegisterSingleton[any], func() {
	var count int
	var observed *fixture.Stateless

	BeforeEach(func() {
		count = 0
		botox.RegisterSingletonSupplier(func() fixture.Stateless {
			count += 1
			return fixture.NewStateless()
		})
	})

	AfterEach(func() {
		botox.Reset()
	})

	Context("with already provided instance", func() {
		var expected *fixture.Stateless

		BeforeEach(func() {
			expected = botox.MustResolve[*fixture.Stateless]()
		})

		It("should return already provided instance", func() {
			observed = botox.MustResolve[*fixture.Stateless]()

			Expect(count).To(Equal(1))
			Expect(observed).To(Equal(expected))
		})
	})
})
