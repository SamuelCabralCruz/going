//go:build test

package pkg2

type SomeStruct struct{}

func (s SomeStruct) Method() string {
	return "some struct from pkg2"
}
