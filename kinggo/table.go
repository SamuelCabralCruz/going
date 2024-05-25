package kinggo

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	"github.com/samber/lo"
	"strings"
)

type tableEntryParameter struct {
	Key   string
	Value any
}

func newTableEntryParameter(key string, value any) tableEntryParameter {
	return tableEntryParameter{Key: key, Value: value}
}

func createTableEntry(parameters ...tableEntryParameter) TableEntry {
	descriptionParts := lo.Map(parameters, func(item tableEntryParameter, _ int) string {
		return fmt.Sprintf("%s: %+v", item.Key, item.Value)
	})
	description := "arguments[" + strings.Join(descriptionParts, ", ") + "]"
	tableEntryParameters := lo.Map(parameters, func(item tableEntryParameter, _ int) any { return item.Value })
	return Entry(description, tableEntryParameters...)
}

func CreateTableEntries(labels []string, cases ...[]any) []TableEntry {
	casesParameters := lo.Map(cases, func(c []any, _ int) []tableEntryParameter {
		return lo.Map(labels, func(label string, index int) tableEntryParameter {
			return newTableEntryParameter(label, c[index])
		})
	})
	return lo.Map(casesParameters, func(caseParameters []tableEntryParameter, _ int) TableEntry {
		return createTableEntry(caseParameters...)
	})
}
