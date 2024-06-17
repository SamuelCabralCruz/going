//go:build test

package fixture

import (
	"errors"
	"github.com/SamuelCabralCruz/going/detox"
)

type SomeStruct struct {
	field1 string
	field2 int
}

type Interface1 interface {
	NoArgNoReturn()
	NoArgSingleReturn() string
	NoArgMultipleReturns() (chan<- string, error, bool)
	SingleArgNoReturn(arg1 string)
	SingleArgSingleReturn(arg1 string) string
	SingleArgMultipleReturns(arg1 string) (string, int)
	MultipleArgsNoReturn(arg1 int, arg2 bool, arg3 []byte)
	MultipleArgsSingleReturn(arg1 int, arg2 bool) string
	MultipleArgsMultipleReturns(arg1 float64, arg2 uint8) (SomeStruct, bool)
}

type Implementation1 struct{}

var _ Interface1 = Implementation1{}

func (i Implementation1) NoArgNoReturn() {}

func (i Implementation1) NoArgSingleReturn() string {
	return "ret1"
}

func (i Implementation1) NoArgMultipleReturns() (chan<- string, error, bool) {
	return make(chan string), errors.New("ret2"), true
}

func (i Implementation1) SingleArgNoReturn(_ string) {}

func (i Implementation1) SingleArgSingleReturn(_ string) string {
	return "ret1"
}

func (i Implementation1) SingleArgMultipleReturns(_ string) (string, int) {
	return "ret1", 2
}

func (i Implementation1) MultipleArgsNoReturn(_ int, _ bool, _ []byte) {}

func (i Implementation1) MultipleArgsSingleReturn(_ int, _ bool) string {
	return "ret1"
}

func (i Implementation1) MultipleArgsMultipleReturns(_ float64, _ uint8) (SomeStruct, bool) {
	return SomeStruct{field1: "field1", field2: 2}, false
}

func NewInterface1Mock() Interface1Mock {
	return Interface1Mock{detox.New[Interface1]()}
}

type Interface1Mock struct {
	*detox.Detox[Interface1]
}

var _ Interface1 = Interface1Mock{}

func (i Interface1Mock) NoArgNoReturn() {
	detox.When(i.Detox, i.NoArgNoReturn).ResolveForArgs()()
}

func (i Interface1Mock) NoArgSingleReturn() string {
	return detox.When(i.Detox, i.NoArgSingleReturn).ResolveForArgs()()
}

func (i Interface1Mock) NoArgMultipleReturns() (chan<- string, error, bool) {
	return detox.When(i.Detox, i.NoArgMultipleReturns).ResolveForArgs()()
}

func (i Interface1Mock) SingleArgNoReturn(arg1 string) {
	detox.When(i.Detox, i.SingleArgNoReturn).ResolveForArgs(arg1)(arg1)
}

func (i Interface1Mock) SingleArgSingleReturn(arg1 string) string {
	return detox.When(i.Detox, i.SingleArgSingleReturn).ResolveForArgs(arg1)(arg1)
}

func (i Interface1Mock) SingleArgMultipleReturns(arg1 string) (string, int) {
	return detox.When(i.Detox, i.SingleArgMultipleReturns).ResolveForArgs(arg1)(arg1)
}

func (i Interface1Mock) MultipleArgsNoReturn(arg1 int, arg2 bool, arg3 []byte) {
	detox.When(i.Detox, i.MultipleArgsNoReturn).ResolveForArgs(arg1, arg2, arg3)(arg1, arg2, arg3)
}

func (i Interface1Mock) MultipleArgsSingleReturn(arg1 int, arg2 bool) string {
	return detox.When(i.Detox, i.MultipleArgsSingleReturn).ResolveForArgs(arg1, arg2)(arg1, arg2)
}

func (i Interface1Mock) MultipleArgsMultipleReturns(arg1 float64, arg2 uint8) (SomeStruct, bool) {
	return detox.When(i.Detox, i.MultipleArgsMultipleReturns).ResolveForArgs(arg1, arg2)(arg1, arg2)
}
