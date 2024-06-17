package trust

import (
	"github.com/SamuelCabralCruz/going/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/SamuelCabralCruz/going/xpctd"
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
