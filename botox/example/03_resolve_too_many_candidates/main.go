package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// Similarly, in case more than one dependency provider are registered, an
	// error will also be returned.

	// REGISTRATION
	botox.RegisterSupplier(example.NewSomeStruct)
	botox.RegisterSupplier(example.NewSomeStruct)
	botox.RegisterSupplier(example.NewSomeStruct)
	// RESOLUTION
	_, err := botox.Resolve[example.SomeStruct]()
	fmt.Println(err.Error())

	// Side Note:
	// For those that have not seen the link between a supplier and a no argument
	// constructor yet, this is what we tried to show here.
	// To use Botox to its full potential, we strongly recommend establishing
	// the standard `AllArgsConstructor` for tests and `NoArgsConstructor` for
	// dependency injection.
	// The `NoArgsConstructor` should invoke the `AllArgsConstructor` after resolving
	// every dependency from Botox.
	// The `AllArgsConstructor` should do simple instantiation (associating
	// arguments to struct fields) or any logic other than resolving dependencies.
}
