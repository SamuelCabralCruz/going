//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// Botox also supports registering singletons.

	// Registering a singleton injection token is different in two ways.

	// First, we will automatically take a reference (pointer) on the value
	// returned by the dependency provider.

	// Second, we will only execute the provider once.
	// No matter the return (success or error) of the dependency provider, the
	// result will be "persisted" and returned immediately for any later
	// resolution.

	// All registration functions have their singleton equivalent.
	// `botox.RegisterInstance` -> `botox.RegisterSingletonInstance`
	// `botox.RegisterSupplier` -> `botox.RegisterSingletonSupplier`
	// `botox.RegisterProducer` -> `botox.RegisterSingletonProducer`

	// REGISTRATION
	botox.RegisterSingletonSupplier(func() example.SomeStruct {
		fmt.Println("singleton supplier invoked")
		return example.SomeStruct{}
	})
	// RESOLUTION
	_, err := botox.Resolve[example.SomeStruct]() // -> should return an error
	fmt.Println(err.Error())
	botox.MustResolve[*example.SomeStruct]() // -> should execute the provider
	botox.MustResolve[*example.SomeStruct]() // -> should not execute the provider
	botox.MustResolve[*example.SomeStruct]() // -> should not execute the provider
	botox.MustResolve[*example.SomeStruct]() // -> should not execute the provider
	botox.MustResolve[*example.SomeStruct]() // -> should not execute the provider

	// Side Note:
	// Please observe the resolution of a pointer on `example.SomeStruct` even
	// if the registered supplier returns an `example.SomeStruct` value.
}
