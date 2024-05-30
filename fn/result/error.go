package result

import "github.com/SamuelCabralCruz/went/roar"

type FilteredValueError struct {
	roar.Roar[FilteredValueError]
}

func newFilteredValueError(value any) FilteredValueError {
	return FilteredValueError{
		Roar: roar.New[FilteredValueError](
			"value does not satisfy predicate",
			roar.WithField("value", value)),
	}
}
