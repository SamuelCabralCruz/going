//go:build test

package phi_test

import (
	. "github.com/SamuelCabralCruz/going/kinggo"
	"github.com/SamuelCabralCruz/going/phi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(phi.HasStructField[any], func() {
	var fieldName string
	var observed bool

	act := func() {
		observed = phi.HasStructField[phi.CustomStruct](fieldName)
	}

	DescribeTable("should detect presence of field", func(tableFieldName string, expected bool) {
		fieldName = tableFieldName

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries([]string{"fieldName", "expected"},
			[]any{"non existing field name", false},
			[]any{"foo", false},
			[]any{"bar", false},
			[]any{"Embedded", true},
			[]any{"nonExportedEmbedded", true},
			[]any{"ExportedStruct", true},
			[]any{"nonExportedStruct", true},
			[]any{"ExportedNonStruct", true},
			[]any{"nonExportedNonStruct", true},
		),
	)
})

var _ = DescribeFunction(phi.HasEmbeddedStructField[any], func() {
	var fieldName string
	var observed bool

	act := func() {
		observed = phi.HasEmbeddedStructField[phi.CustomStruct](fieldName)
	}

	DescribeTable("should detect presence of embedded field", func(tableFieldName string, expected bool) {
		fieldName = tableFieldName

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries([]string{"fieldName", "expected"},
			[]any{"non existing field name", false},
			[]any{"foo", false},
			[]any{"bar", false},
			[]any{"Embedded", true},
			[]any{"nonExportedEmbedded", true},
			[]any{"ExportedStruct", false},
			[]any{"nonExportedStruct", false},
			[]any{"ExportedNonStruct", false},
			[]any{"nonExportedNonStruct", false},
		),
	)
})
