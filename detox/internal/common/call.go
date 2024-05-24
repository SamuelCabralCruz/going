package common

import (
	"github.com/SamuelCabralCruz/went/fn"
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
	return fn.Copy(c.args)
}

func (c Call) EqualTo(other Call) bool {
	return reflect.DeepEqual(c.args, other.args)
}
