//go:build example

package main

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// The Botox API is pretty minimalist while still offering plenty of leeway
	// to match all your desires. 🍑🍒🤞

	// To register a provider, you can use the following functions:
	// - `botox.RegisterInstance`
	// - `botox.RegisterSupplier`
	// - `botox.RegisterProducer`

	// Ultimately, all those functions will end up registering a producer.
	// In the following, we might use the term dependency provider interchangeably
	// with producer.
	// In other words, a dependency provider is simply a function taking no argument
	// and returning an instance and/or an error.

	// REGISTRATION
	botox.RegisterInstance(example.SomeStruct{})
	botox.RegisterSupplier(func() example.SomeStruct {
		return example.SomeStruct{}
	})
	botox.RegisterProducer(func() (example.SomeStruct, error) {
		return example.SomeStruct{}, nil
	})
}
