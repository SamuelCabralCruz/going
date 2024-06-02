package thong

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

func Indent(text string, separator string, indentation string) string {
	return strings.Join(
		indentParts(
			indentation,
			strings.Split(text, separator)),
		separator)
}

func IndentParts(indentation string, parts []string) string {
	return strings.Join(indentParts(indentation, parts), "\n")
}

func indentParts(indentation string, parts []string) []string {
	return lo.Map(parts, func(part string, _ int) string {
		return fmt.Sprintf("%s%s", indentation, part)
	})
}

func Surround(text string, chars string) string {
	return fmt.Sprintf("%s%s%s", chars, text, chars)
}

func Quote(text string) string {
	return Surround(text, `"`)
}
