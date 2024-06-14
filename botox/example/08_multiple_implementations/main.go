//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// A common use case when dealing with dependency injection is the ability
	// to register multiple implementations of the same interface.

	// To achieve this using Botox, you must make sure to register your instances
	// on the interface, not their actual type.

	// REGISTRATION
	botox.RegisterInstance(example.SomeImplementation1{})
	botox.RegisterInstance(example.SomeImplementation2{})
	botox.RegisterInstance(example.SomeImplementation3{})
	// RESOLUTION
	instances, err := botox.ResolveAll[example.SomeInterface]()
	fmt.Println(fmt.Sprintf("nb instances: %d", len(instances)))
	fmt.Println(err.Error())

	// vs
	botox.Reset()

	// REGISTRATION
	botox.RegisterInstance[example.SomeInterface](example.SomeImplementation1{})
	botox.RegisterInstance[example.SomeInterface](example.SomeImplementation2{})
	botox.RegisterInstance[example.SomeInterface](example.SomeImplementation3{})
	// RESOLUTION
	instances, err = botox.ResolveAll[example.SomeInterface]()
	fmt.Println(fmt.Sprintf("nb instances: %d", len(instances)))
	fmt.Println(fmt.Sprintf("error is %T", err))

	// Side Note:
	// Smart cookies 🍪 might have observed the use of `botox.Reset`.
	// Although we don't see an actual need for this feature in production code,
	// we implemented it for testing purposes.
	// Feel free to use it... 🕊️
	// However, keep in mind that it resets the entire container...
	// Given you would like to unregister all tokens for a given type, you can have
	// a look at `botox.Unregister`. 🧹
}
