//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/going/botox"
	"github.com/SamuelCabralCruz/going/botox/example/00_boilerplate"
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

	// The quicker among you might have realized that registering a pointer on
	// a value is equivalent to registering the singleton of the underlying value.

	// In the following table, we disclose the association between the registration
	// and the resolution depending on the Botox function used.

	// | Registered         | Function                        | Resolved |
	// | ------------------ | ------------------------------- | -------- |
	// | T                  | botox.RegisterInstance          | T        |
	// | func() T           | botox.RegisterSupplier          | T        |
	// | func() (T, error)  | botox.RegisterProducer          | T        |
	// | *T                 | botox.RegisterInstance          | *T       |
	// | func() *T          | botox.RegisterSupplier          | *T       |
	// | func() (*T, error) | botox.RegisterProducer          | *T       |
	// | T                  | botox.RegisterSingletonInstance | *T       |
	// | func() T           | botox.RegisterSingletonSupplier | *T       |
	// | func() (T, error)  | botox.RegisterSingletonProducer | *T       |
	// | *T                 | botox.RegisterSingletonInstance | **T      |
	// | func() *T          | botox.RegisterSingletonSupplier | **T      |
	// | func() (*T, error) | botox.RegisterSingletonProducer | **T      |
}
