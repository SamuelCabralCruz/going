//go:build test

package roar_test

import (
	"github.com/SamuelCabralCruz/went/roar"
	. "github.com/SamuelCabralCruz/went/testing"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[roar.Roar[any]](func() {
	FIt("should work", func() {
		Expect(true).To(BeTrue())
	})
})
