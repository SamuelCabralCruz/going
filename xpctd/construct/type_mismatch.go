package construct

import (
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/SamuelCabralCruz/going/xpctd"
)

func TypeMismatchReporter[T any]() xpctd.Reporter[any] {
	return xpctd.Actual[any]().
		ToBeOfType(phi.BaseTypeName[T]()).
		ButWasOfType()
}
