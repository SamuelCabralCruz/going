//go:build test

package usage_test

import (
	"errors"
	"github.com/SamuelCabralCruz/going/botox"
	"github.com/SamuelCabralCruz/going/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.MustResolve[any], func() {
	act := func() {
		botox.MustResolve[fixture.Stateless]()
	}

	AfterEach(func() {
		botox.Reset()
	})

	Context("with producer returning error registration", func() {
		var producerError error

		BeforeEach(func() {
			producerError = errors.New("something went wrong")
			botox.RegisterProducer(func() (fixture.Stateless, error) {
				return fixture.Stateless{}, producerError
			})
		})

		It("should panic", func() {
			Expect(act).To(PanicWith(producerError))
		})
	})
})
