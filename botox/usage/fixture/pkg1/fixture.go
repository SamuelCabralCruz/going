//go:build test

package pkg1

type SomeStruct struct{}

func (s SomeStruct) Method() string {
	return "some struct from pkg1"
}
