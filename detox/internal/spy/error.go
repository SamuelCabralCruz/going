package spy

import (
	"github.com/SamuelCabralCruz/went/detox/internal/common"
	"github.com/SamuelCabralCruz/went/roar"
)

type InvalidCallIndexError struct {
	roar.Roar[InvalidCallIndexError]
}

func newInvalidCallIndexError(info common.MockedMethodInfo, received int, nbCalls int) InvalidCallIndexError {
	return InvalidCallIndexError{
		roar.New[InvalidCallIndexError](
			"invalid call index provided",
			roar.WithField("interface", info.Interface()),
			roar.WithField("method", info.Method()),
			roar.WithField("reference", info.Reference()),
			roar.WithField("received", received),
			roar.WithField("nbCalls", nbCalls))}
}
