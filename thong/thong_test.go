//go:build test

package thong_test

import (
	. "github.com/SamuelCabralCruz/going/kinggo"
	"github.com/SamuelCabralCruz/going/thong"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(thong.Indent, func() {
	var input string
	var separator string
	var indentation string
	var observed string

	act := func() {
		observed = thong.Indent(input, separator, indentation)
	}

	DescribeTable("should indent text properly", func(
		tableInput string,
		tableSeparator string,
		tableIndentation string,
		expected string,
	) {
		input = tableInput
		separator = tableSeparator
		indentation = tableIndentation

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries(
			[]string{"input", "separator", "indentation", "expected"},
			[]any{"a\nb", "\n", "\t", "\ta\n\tb"},
			[]any{"a\nb\nc\td", "\n", "\t\t", "\t\ta\n\t\tb\n\t\tc\td"},
		))
})

var _ = DescribeFunction(thong.IndentParts, func() {
	var indentation string
	var input []string
	var observed string

	act := func() {
		observed = thong.IndentParts(indentation, input)
	}

	DescribeTable("should indent text properly", func(
		tableIndentation string,
		tableInput []string,
		expected string,
	) {
		indentation = tableIndentation
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries(
			[]string{"indentation", "input", "expected"},
			[]any{"\t", []string{"a", "b"}, "\ta\n\tb"},
			[]any{"\t\t", []string{"a", "b", "c\td"}, "\t\ta\n\t\tb\n\t\tc\td"},
		))
})

var _ = DescribeFunction(thong.Surround, func() {
	var input string
	var chars string
	var observed string

	act := func() {
		observed = thong.Surround(input, chars)
	}

	DescribeTable("should surround text properly", func(
		tableInput string,
		tableChars string,
		expected string,
	) {
		input = tableInput
		chars = tableChars

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries(
			[]string{"input", "chars", "expected"},
			[]any{"some text", "-aaa-", "-aaa-some text-aaa-"},
		))
})

var _ = DescribeFunction(thong.Quote, func() {
	var input string
	var observed string

	act := func() {
		observed = thong.Quote(input)
	}

	DescribeTable("should quote text properly", func(
		tableInput string,
		expected string,
	) {
		input = tableInput

		act()

		Expect(observed).To(Equal(expected))
	},
		CreateTableEntries(
			[]string{"input", "expected"},
			[]any{"some text", `"some text"`},
		))
})
