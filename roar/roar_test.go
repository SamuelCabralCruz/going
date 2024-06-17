//go:build test

package roar

import (
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeType[Roar[any]](func() {
	DescribeFunction(Roar[any]{}.Error, func() {
		var cut error
		var observed string

		act := func() {
			observed = cut.Error()
		}

		DescribeTable("should serialize properly", func(tableCut error, expected string) {
			cut = tableCut

			act()

			Expect(observed).To(Equal(expected))
		},
			CreateTableEntries([]string{"cut", "expected"},
				[]any{someErrorBase,
					"someError: base"},
				[]any{someErrorWithStandardErrorCause,
					"someError: with standard error cause - standard error"},
				[]any{someErrorWithMultipleCauses,
					"someError: with multiple causes - last cause"},
				[]any{someErrorWithField,
					"someError: with field [name=value]"},
				[]any{someErrorWithFields,
					"someError: with fields [name1=value1, name2=value2]"},
				[]any{someErrorWithCauseAndFields,
					"someError: with cause and fields [name1=value1, name2=value2] - standard error"},
				[]any{someErrorWithNestedCause,
					"someError: with nested cause - anotherError: nested - standard error"},
				[]any{someErrorWithEverything,
					"someError: with everything [name3=value3, name4=value4] - anotherError: nested [name1=value1, name2=value2] - standard error"},
			),
		)
	})
})
