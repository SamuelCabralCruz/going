package construct

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/gomicron/reporter"
	"github.com/SamuelCabralCruz/went/phi"
)

func TypeMismatchReporter[T any]() reporter.Reporter[any] {
	return reporter.Actual[any]().
		ToBeFormatted(`a "%s"`, phi.BaseTypeName[T]()).
		ButWas(func(actual any) string {
			return fmt.Sprintf(`a "%T"`, actual)
		})
}
