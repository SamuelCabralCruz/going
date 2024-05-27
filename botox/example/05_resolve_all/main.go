//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// In case there are multiple dependency providers to be resolved, you can
	// instead use the `botox.ResolveAll` function which expect to find at least
	// one dependency provider.

	// REGISTRATION
	botox.RegisterSupplier(example.NewSomeStruct)
	botox.RegisterSupplier(example.NewSomeStruct)
	botox.RegisterSupplier(example.NewSomeStruct)
	// RESOLUTION
	instances, err := botox.ResolveAll[example.SomeStruct]()
	fmt.Println(fmt.Sprintf("error is %T", err))
	for _, instance := range instances {
		fmt.Println(instance.Describe())
	}
}
