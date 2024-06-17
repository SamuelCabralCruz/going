package common

import (
	"github.com/SamuelCabralCruz/going/scallop"
	"reflect"
)

type Call struct {
	args []any
}

func NewCall(args ...any) Call {
	return Call{
		args,
	}
}

func (c Call) Args() []any {
	return scallop.Copy(c.args)
}

func (c Call) EqualTo(other Call) bool {
	return reflect.DeepEqual(c.args, other.args)
}
