//go:build test

package phi_test

import (
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/phi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(phi.IsZero, func() {
	var input any
	var observed bool

	act := func() {
		observed = phi.IsZero(input)
	}

	DescribeTable("should detect null value", func(tableInput any, expected bool) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	}, CreateTableEntries([]string{"input", "expected"},
		[]any{func() {}, false},
		[]any{func() { _ = "" }, false},
		[]any{map[string]any{}, true},
		[]any{map[string]any{"a": 1}, false},
		[]any{[]string{}, true},
		[]any{[]string{"a"}, false},
		[]any{phi.Empty[string](), true},
		[]any{"", true},
		[]any{"something", false},
		[]any{0, true},
		[]any{1, false},
		[]any{struct{}{}, true},
		[]any{struct{ a int }{}, true},
		[]any{struct{ a map[string]int }{}, true},
		[]any{struct{ a map[string]int }{a: map[string]int{"b": 2}}, false},
		[]any{struct{ a []int }{}, true},
		[]any{struct{ a []int }{a: []int{1}}, false},
		[]any{struct{ a int }{a: 1}, false},
	))
})

var _ = DescribeFunction(phi.IsTypeOf[any], func() {
	var input any
	var observed bool

	act := func() {
		observed = phi.IsTypeOf[phi.Iam](input)
	}

	DescribeTable("should detect which value implement the specific interface", func(tableInput any, expected bool) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	}, CreateTableEntries([]string{"input", "expected"},
		[]any{phi.IamImplementing{}, true},
		[]any{phi.IamNotImplementing{}, false},
	))
})

var _ = DescribeFunction(phi.IsFunction, func() {
	var input any
	var observed bool

	act := func() {
		observed = phi.IsFunction(input)
	}

	DescribeTable("should detect function", func(tableInput any, expected bool) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	}, CreateTableEntries([]string{"input", "expected"},
		[]any{func() {}, true},
		[]any{func() string { return "" }, true},
		[]any{"", false},
		[]any{1, false},
		[]any{true, false},
		[]any{struct{}{}, false},
		[]any{map[string]any{}, false},
		[]any{[]string{}, false},
	))
})

var _ = DescribeFunction(phi.IsMap, func() {
	var input any
	var observed bool

	act := func() {
		observed = phi.IsMap(input)
	}

	DescribeTable("should detect map", func(tableInput any, expected bool) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	}, CreateTableEntries([]string{"input", "expected"},
		[]any{func() {}, false},
		[]any{func() string { return "" }, false},
		[]any{"", false},
		[]any{1, false},
		[]any{true, false},
		[]any{struct{}{}, false},
		[]any{map[string]any{}, true},
		[]any{[]string{}, false},
	))
})

var _ = DescribeFunction(phi.IsSlice, func() {
	var input any
	var observed bool

	act := func() {
		observed = phi.IsSlice(input)
	}

	DescribeTable("should detect slice", func(tableInput any, expected bool) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	}, CreateTableEntries([]string{"input", "expected"},
		[]any{func() {}, false},
		[]any{func() string { return "" }, false},
		[]any{"", false},
		[]any{1, false},
		[]any{true, false},
		[]any{struct{}{}, false},
		[]any{map[string]any{}, false},
		[]any{[]string{}, true},
	))
})
