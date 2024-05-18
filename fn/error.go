package fn

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/roar"
	"github.com/samber/lo"
	"strings"
)

type AggregatedError struct {
	roar.Roar[AggregatedError]
}

func NewAggregatedError(errs ...error) AggregatedError {
	return AggregatedError{
		roar.New[AggregatedError](
			fmt.Sprintf("multiple errors occurred - [%s]",
				strings.Join(lo.Map(errs, func(err error, _ int) string {
					return err.Error()
				}), ", "))),
	}
}
