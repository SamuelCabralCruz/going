//go:build test

package xpctd_test

import (
	"fmt"
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/xpctd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = DescribeType[xpctd.Reporter[any]](func() {
	var actual string
	var cut xpctd.Reporter[string]

	BeforeEach(func() {
		actual = "some actual value"
	})

	cases := CreateTableEntries([]string{"case", "reporter", "expected"},
		[]any{"default", xpctd.New[string](), "Expected \"some actual value\" to match"},
		[]any{
			"computed",
			xpctd.Computed[string](func(actual string) string {
				return fmt.Sprintf("some computed value - %s", strings.ToTitle(actual))
			}),
			"Expected some computed value - SOME ACTUAL VALUE to match"},
		[]any{"formatted", xpctd.Formatted[string]("%s %d", "formatted", 1234), "Expected formatted 1234 to match"},
		[]any{"the", xpctd.The[string]("description"), "Expected the description to match"},
		[]any{"type", xpctd.Type[string](), "Expected string to match"},
		[]any{"value", xpctd.Value[string](), "Expected value to match"},
		[]any{"actual", xpctd.Actual[string](), "Expected \"some actual value\" to match"},
		[]any{"actual negative", xpctd.Actual[string]().Negative(), "Expected \"some actual value\" not to match"},
		[]any{
			"actual to",
			xpctd.Actual[string]().To(func(actual string) string {
				return fmt.Sprintf("be a %T of length %d", actual, len(actual))
			}),
			"Expected \"some actual value\" to be a string of length 17"},
		[]any{"actual to formatted", xpctd.Actual[string]().ToFormatted("be a %s of length %d", "string", 12), "Expected \"some actual value\" to be a string of length 12"},
		[]any{
			"actual to be",
			xpctd.Actual[string]().ToBe(func(actual string) string {
				return fmt.Sprintf("a %T of length %d", actual, len(actual))
			}),
			"Expected \"some actual value\" to be a string of length 17"},
		[]any{"actual to be formatted", xpctd.Actual[string]().ToBeFormatted("a %s of length %d", "string", 12), "Expected \"some actual value\" to be a string of length 12"},
		[]any{"actual to be a", xpctd.Actual[string]().ToBeA("string of length 12"), "Expected \"some actual value\" to be a string of length 12"},
		[]any{"actual to be of type", xpctd.Actual[string]().ToBeOfType("some type name"), "Expected \"some actual value\" to be of type \"some type name\""},
		[]any{
			"actual to have",
			xpctd.Actual[string]().ToHave(func(actual string) string {
				return fmt.Sprintf("length %d", len(actual))
			}),
			"Expected \"some actual value\" to have length 17"},
		[]any{"actual to have formatted", xpctd.Actual[string]().ToHaveFormatted("length %d", 12), "Expected \"some actual value\" to have length 12"},
		[]any{"actual but", xpctd.Actual[string]().But(func(actual string) string {
			return fmt.Sprintf("did not (length: %d)", len(actual))
		}), "Expected \"some actual value\" to match, but did not (length: 17)"},
		[]any{"actual but formatted", xpctd.Actual[string]().ButFormatted("did not (length: 12)"), "Expected \"some actual value\" to match, but did not (length: 12)"},
		[]any{"actual but received", xpctd.Actual[string]().ButReceived(func(actual string) string {
			return fmt.Sprintf("%T of length %d", actual, len(actual))
		}), "Expected \"some actual value\" to match, but received string of length 17"},
		[]any{"actual but received formatted", xpctd.Actual[string]().ButReceivedFormatted("string of length %d", 12), "Expected \"some actual value\" to match, but received string of length 12"},
		[]any{"actual but was", xpctd.Actual[string]().ButWas(func(actual string) string {
			return fmt.Sprintf("%T of length %d", actual, len(actual))
		}), "Expected \"some actual value\" to match, but was string of length 17"},
		[]any{"actual but was formatted", xpctd.Actual[string]().ButWasFormatted("string of length %d", 12), "Expected \"some actual value\" to match, but was string of length 12"},
		[]any{"actual but was a", xpctd.Actual[string]().ButWasA("string of length 12"), "Expected \"some actual value\" to match, but was a string of length 12"},
		[]any{"actual but was of type", xpctd.Actual[string]().ButWasOfType(), "Expected \"some actual value\" to match, but was of type \"string\""},
		[]any{"actual but had", xpctd.Actual[string]().ButHad(func(actual string) string {
			return fmt.Sprintf("length %d", len(actual))
		}), "Expected \"some actual value\" to match, but had length 17"},
		[]any{"actual but was formatted", xpctd.Actual[string]().ButHadFormatted("length %d", 12), "Expected \"some actual value\" to match, but had length 12"},
	)

	DescribeFunction(xpctd.Reporter[any].Report, func() {
		var observed string

		act := func() {
			observed = cut.Report(actual)
		}

		DescribeTable("should report properly", func(_ string, reporter xpctd.Reporter[string], expected string) {
			cut = reporter

			act()

			Expect(observed).To(Equal(expected))
		}, cases)
	})

	DescribeFunction(xpctd.Reporter[any].Error, func() {
		var observed error

		act := func() {
			observed = cut.Error(actual)
		}

		DescribeTable("should return error with corresponding report", func(_ string, reporter xpctd.Reporter[string], expected string) {
			cut = reporter

			act()

			Expect(observed).To(BeAssignableToTypeOf(xpctd.ExpectationError{}))
			Expect(observed.Error()).To(ContainSubstring(expected))
		}, cases)
	})
})
