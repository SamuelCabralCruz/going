package construct

import (
	"github.com/SamuelCabralCruz/went/phi"
	reporter2 "github.com/SamuelCabralCruz/went/xpctd"
)

func TypeMismatchReporter[T any]() reporter2.Reporter[any] {
	return reporter2.Actual[any]().
		ToBeOfType(phi.Empty[T]()).
		ButWasOfType()
}
