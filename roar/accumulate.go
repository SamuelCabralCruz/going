package roar

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/samber/mo"
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

func Accumulate[T any](results ...mo.Result[T]) mo.Result[[]T] {
	var acc []T
	var errors []error
	lo.ForEach(results,
		func(result mo.Result[T], _ int) {
			if result.IsOk() {
				acc = append(acc, result.MustGet())
			} else {
				errors = append(errors, result.Error())
			}
		})
	if len(errors) > 0 {
		return mo.Err[[]T](NewAggregatedError(errors...))
	}
	return mo.Ok(acc)
}
