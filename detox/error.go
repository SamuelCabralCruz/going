package detox

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/roar"
)

type InterfaceMethodMismatchError struct {
	roar.Roar[InterfaceMethodMismatchError]
}

func newInterfaceMethodMismatchError(info common.MockedMethodInfo, reason string) InterfaceMethodMismatchError {
	return InterfaceMethodMismatchError{
		roar.New[InterfaceMethodMismatchError](
			"specified method does not belong to the mocked interface",
			roar.WithField("interface", info.Interface()),
			roar.WithField("method", info.Method()),
			roar.WithField("reason", reason),
		),
	}
}
