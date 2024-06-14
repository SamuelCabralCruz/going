package construct

import (
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/SamuelCabralCruz/went/xpctd"
)

func TypeMismatchReporter[T any]() xpctd.Reporter[any] {
	return xpctd.Actual[any]().
		ToBeOfType(phi.BaseTypeName[T]()).
		ButWasOfType()
}
