//go:build test

package fixture

type SomeInterface interface {
	Method() string
}

type SomeImplementation1 struct{}

var _ SomeInterface = SomeImplementation1{}

func (s SomeImplementation1) Method() string {
	return "implementation 1"
}

type SomeImplementation2 struct{}

var _ SomeInterface = SomeImplementation2{}

func (s SomeImplementation2) Method() string {
	return "implementation 2"
}

type SomeImplementation3 struct{}

var _ SomeInterface = SomeImplementation3{}

func (s SomeImplementation3) Method() string {
	return "implementation 3"
}
