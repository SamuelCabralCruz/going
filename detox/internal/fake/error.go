package fake

import (
	"github.com/SamuelCabralCruz/going/detox/internal/common"
	"github.com/SamuelCabralCruz/going/roar"
)

type MissingImplementationError struct {
	roar.Roar[MissingImplementationError]
}

func newMissingImplementationError(info common.MockedMethodInfo, call common.Call) MissingImplementationError {
	return MissingImplementationError{
		roar.New[MissingImplementationError](
			"no implementation has been registered",
			roar.WithField("interface", info.Interface()),
			roar.WithField("method", info.Method()),
			roar.WithField("arguments", call.Args()),
			roar.WithField("reference", info.Reference()),
		),
	}
}
