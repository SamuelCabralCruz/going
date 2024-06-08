//go:build test

package phi

type SomeInterface interface {
	A()
	B(string)
	C() string
}

type SomeImplementation struct {
}

func (s SomeImplementation) A() {}

func (s SomeImplementation) B(_ string) {}

func (s SomeImplementation) C() string { return "" }

var _ SomeInterface = SomeImplementation{}

type AnotherImplementation struct{}

func (a AnotherImplementation) A(_ string) {}

func (a AnotherImplementation) B(_ string) {}

func (a AnotherImplementation) D(_ int) string { return "" }

type Embedded struct{}

type nonExportedEmbedded struct{}

type CustomStruct struct {
	Embedded
	nonExportedEmbedded
	ExportedStruct       struct{}
	nonExportedStruct    struct{}
	ExportedNonStruct    string
	nonExportedNonStruct string
}

type AnonymousStruct = struct{ field string }

type Iam interface {
	groot() string
}

type IamImplementing struct{}

func (_ IamImplementing) groot() string {
	return "groot"
}

type IamNotImplementing struct{}

func (_ IamNotImplementing) Speak() string {
	return "speak"
}

func CustomFunction() {}

func GenericFunction[_ any]() {}

var AnonymousFunction = func() {}
