//go:build test

package phi

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

func GenericFunction[T any]() {}

var AnonymousFunction = func() {}
