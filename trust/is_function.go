package trust

import (
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/SamuelCabralCruz/went/xpctd"
)

func AssertIsFunction(value any) (any, error) {
	if phi.IsFunction(value) {
		return value, nil
	}
	return nil, xpctd.Value[any]().
		ToBeA("function").
		ButWasOfType().
		Error(value)
}

func IsFunction(value any) (any, bool) {
	return assertion.ToValidation(AssertIsFunction(value))
}
