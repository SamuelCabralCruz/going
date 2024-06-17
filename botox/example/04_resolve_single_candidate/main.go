//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/going/botox"
	"github.com/SamuelCabralCruz/going/botox/example/00_boilerplate"
)

func main() {
	// When a single dependency provider has been registered, the corresponding
	// instance will be returned.

	// REGISTRATION
	botox.RegisterSupplier(example.NewSomeStruct)
	// RESOLUTION
	instance, err := botox.Resolve[example.SomeStruct]()
	fmt.Println(fmt.Sprintf("error is %T", err))
	fmt.Println(instance.Describe())
}
