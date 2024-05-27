//go:build example

package main

import (
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// Dealing with errors is not always pleasant.
	// This is especially true in the initialization loop of your application. 💩
	// This is the reason why we decided to expose the method `botox.MustResolve`.

	// The only difference with `botox.Resolve` is the fact that this method
	// will panic when an error would normally be returned.

	// Using this method, you might be able to simplify your application logic
	// instead of having to deal with errors returning functions when there is
	// no foreseeable recovery possible.

	// REGISTRATION
	// RESOLUTION
	botox.MustResolve[example.SomeStruct]() // -> should panic
}
