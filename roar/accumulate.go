package roar

import (
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type AggregatedError struct {
	Roar[AggregatedError]
}

func NewAggregatedError(errs ...error) AggregatedError {
	return AggregatedError{
		New[AggregatedError](
			fmt.Sprintf("multiple errors occurred - [%s]",
				strings.Join(lo.Map(errs, func(err error, _ int) string {
					return err.Error()
				}), ", "))),
	}
}

func Combine(err1 error, err2 error) error {
	if err1 == nil {
		return err2
	}
	// TODO: rename -> Combined error maybe
	return NewAggregatedError(err1, err2)
}
