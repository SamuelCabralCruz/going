//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// To resolve a dependency, you can use the `botox.Resolve` function.

	// This function expects to find exactly one dependency provider (registration)
	// for the specified type.

	// In case no dependency provider is found, it will return an error alongside
	// a zero value of the desired type.

	// REGISTRATION
	// RESOLUTION
	_, err := botox.Resolve[example.SomeStruct]()
	fmt.Println(err.Error())
}
