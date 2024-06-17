package phi_test

import (
	. "github.com/SamuelCabralCruz/going/kinggo"
	"github.com/SamuelCabralCruz/going/phi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = DescribeFunction(phi.GetInterfaceMethodByName[any], func() {
	var input string
	var observedValue reflect.Method
	var observedFound bool

	act := func() {
		observedValue, observedFound = phi.GetInterfaceMethodByName[phi.SomeInterface](input)
	}

	Context("with existing method", func() {
		DescribeTable("should return corresponding method", func(
			tableInput string,
			expected any,
		) {
			input = tableInput

			act()

			Expect(observedFound).To(BeTrue())
			Expect(observedValue).NotTo(BeZero())
			Expect(observedValue.Name).To(Equal(input))
			Expect(observedValue.Type).To(Equal(phi.Value(expected).Type()))
		},
			CreateTableEntries([]string{"input", "expected"},
				[]any{"A", phi.SomeImplementation{}.A},
				[]any{"B", phi.SomeImplementation{}.B},
				[]any{"C", phi.SomeImplementation{}.C},
			),
		)
	})

	Context("with non existing method", func() {
		DescribeTable("should return no method", func(
			tableInput string,
		) {
			input = tableInput

			act()

			Expect(observedFound).To(BeFalse())
			Expect(observedValue).To(BeZero())
		},
			CreateTableEntries([]string{"input"},
				[]any{"a"},
				[]any{"Aa"},
			),
		)
	})
})

var _ = DescribeFunction(phi.GetMatchingInterfaceMethod[any, any], func() {
	var input any
	var observedValue reflect.Method
	var observedError error

	act := func() {
		observedValue, observedError = phi.GetMatchingInterfaceMethod[phi.SomeInterface, any](input)
	}

	Context("with existing method", func() {
		DescribeTable("should return corresponding method", func(
			tableInput any,
		) {
			input = tableInput

			act()

			Expect(observedError).To(BeNil())
			Expect(observedValue).NotTo(BeZero())
			Expect(observedValue.Name).To(Equal(phi.FunctionName(input)))
			Expect(observedValue.Type).To(Equal(phi.Value(input).Type()))
		},
			CreateTableEntries([]string{"input"},
				[]any{phi.SomeImplementation{}.A},
				[]any{phi.SomeImplementation{}.B},
				[]any{phi.SomeImplementation{}.C},
				[]any{phi.AnotherImplementation{}.B},
			),
		)
	})

	Context("with non existing method", func() {
		DescribeTable("should return error", func(
			tableInput any,
			expectedError string,
		) {
			input = tableInput

			act()

			Expect(observedError).NotTo(BeNil())
			Expect(observedError.Error()).To(Equal(expectedError))
			Expect(observedValue).To(BeZero())
		},
			CreateTableEntries([]string{"input", "expectedError"},
				[]any{phi.AnotherImplementation{}.A, "interface method type \"func()\" does not match provided method type \"func(string)\""},
				[]any{phi.AnotherImplementation{}.D, "interface \"SomeInterface\" does not have a method named \"D\""},
			),
		)
	})
})
