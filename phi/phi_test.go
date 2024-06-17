//go:build test

package phi_test

import (
	. "github.com/SamuelCabralCruz/going/kinggo"
	"github.com/SamuelCabralCruz/going/phi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(phi.UniqueIdentifier[any], func() {
	DescribeTable("should compute proper unique identifier", func(act func() string, expected string) {
		observed := act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries([]string{"act", "expected"},
			[]any{phi.UniqueIdentifier[phi.CustomStruct], "github.com/SamuelCabralCruz/going/phi.phi.CustomStruct[struct]"},
			[]any{phi.UniqueIdentifier[phi.AnonymousStruct], ".struct { field string }[struct]"},
			[]any{phi.UniqueIdentifier[string], ".string[string]"},
			[]any{phi.UniqueIdentifier[bool], ".bool[bool]"},
			[]any{phi.UniqueIdentifier[int], ".int[int]"},
			[]any{phi.UniqueIdentifier[float64], ".float64[float64]"},
			[]any{phi.UniqueIdentifier[float32], ".float32[float32]"},
			[]any{phi.UniqueIdentifier[error], ".error[interface]"},
			[]any{phi.UniqueIdentifier[phi.Iam], "github.com/SamuelCabralCruz/going/phi.phi.Iam[interface]"},
		),
	)
})
