//go:build example

package example

import (
	"fmt"
	"time"
)

// Using Botox, you can leverage an injection token pattern to build an ambient
// container containing all your instances, suppliers, and/or producers.

// In short, it is a dependency injection (DI) framework using golang generics.

// Before demonstrating all its capacities, let's first define some fixtures we
// will reuse across the examples.

func NewSomeStruct() SomeStruct {
	return SomeStruct{
		timestamp: time.Now().Nanosecond(),
	}
}

type SomeStruct struct {
	timestamp int
}

func (s SomeStruct) Describe() string {
	return fmt.Sprintf("some struct - %d", s.timestamp)
}

type SomeInterface interface {
	Method(string) string
}

type SomeImplementation1 struct{}

var _ SomeInterface = SomeImplementation1{}

func (s SomeImplementation1) Method(arg string) string {
	return fmt.Sprintf("implementation 1 - %s", arg)
}

type SomeImplementation2 struct{}

var _ SomeInterface = SomeImplementation2{}

func (s SomeImplementation2) Method(arg string) string {
	return fmt.Sprintf("implementation 2 - %s", arg)
}

type SomeImplementation3 struct{}

var _ SomeInterface = SomeImplementation3{}

func (s SomeImplementation3) Method(arg string) string {
	return fmt.Sprintf("implementation 3 - %s", arg)
}
