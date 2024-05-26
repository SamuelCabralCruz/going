//go:build test

package matcher_test

import (
	"errors"
	"github.com/SamuelCabralCruz/went/detox"
	"github.com/SamuelCabralCruz/went/detox/example2/pkg"
	"github.com/SamuelCabralCruz/went/detox/matcher"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[detox.Detox[any]](func() {
	mock := pkg.NewSomeMockClass()
	mockedHello := detox.When(mock.Detox, mock.Hello)

	BeforeEach(func() {
		mock.Default(pkg.Impl{})
		mockedHello.Call(func(s string) (string, error) {
			return "", errors.New("using mock to stub error")
		})
	})

	AfterEach(func() {
		mock.Reset()
	})

	DescribeFunction(matcher.HaveBeenCalled, func() {
		It("should determine whether mock has been called", func() {
			_, _ = mock.Hello("some arg")

			Expect(mockedHello).To(matcher.HaveBeenCalled())
		})
	})
})
