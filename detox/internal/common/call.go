package common

import "reflect"

type Call struct {
	args []any
}

func NewCall(args ...any) Call {
	return Call{
		args,
	}
}

func (c Call) EqualTo(other Call) bool {
	return reflect.DeepEqual(c.args, other.args)
}
