//go:build test

package phi_test

import (
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/phi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(phi.TypeName[any], func() {
	DescribeTable("should compute proper name", func(act func() string, expected string) {
		observed := act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries([]string{"act", "expected"},
			[]any{phi.TypeName[phi.CustomStruct], "CustomStruct"},
			[]any{phi.TypeName[phi.AnonymousStruct], ""},
			[]any{phi.TypeName[string], "string"},
			[]any{phi.TypeName[bool], "bool"},
			[]any{phi.TypeName[int], "int"},
			[]any{phi.TypeName[float64], "float64"},
			[]any{phi.TypeName[float32], "float32"},
			[]any{phi.TypeName[error], "error"},
			[]any{phi.TypeName[phi.Iam], "Iam"},
		),
	)
})

var _ = DescribeFunction(phi.FunctionName, func() {
	var input any
	var observed string

	act := func() {
		observed = phi.FunctionName(input)
	}

	DescribeTable("should compute proper name", func(tableInput any, expected string) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries([]string{"input", "expected"},
			[]any{func() {}, "func2.3"},
			[]any{phi.AnonymousFunction, "func1"},
			[]any{phi.CustomFunction, "CustomFunction"},
			[]any{phi.GenericFunction[any], "GenericFunction"},
		),
	)
})
