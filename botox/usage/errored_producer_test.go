//go:build test

package usage_test

import (
	"errors"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/usage/fixture"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(botox.MustResolve[any], func() {
	act := func() {
		botox.MustResolve[fixture.Stateless]()
	}

	AfterEach(func() {
		botox.Clear()
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
