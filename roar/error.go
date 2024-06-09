package roar

import (
	"fmt"
	"github.com/samber/lo"
)

type AggregatedError struct {
	Roar[AggregatedError]
	errs []error
}

func NewAggregatedError(errs ...error) AggregatedError {
	return AggregatedError{
		New[AggregatedError](
			"multiple errors occurred",
			lo.Map(errs, func(err error, index int) Option {
				return WithField(fmt.Sprintf("[%d]", index), err.Error())
			})...),
		errs,
	}
}

func (a AggregatedError) Errors() []error {
	return a.errs
}
