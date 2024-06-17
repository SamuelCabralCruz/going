//go:build example

package main

import (
	"github.com/SamuelCabralCruz/going/botox"
	"github.com/SamuelCabralCruz/going/botox/example/00_boilerplate"
)

func main() {
	// Similarly, we did the equivalent for `botox.ResolveAll`. 🙇

	// REGISTRATION
	// RESOLUTION
	botox.MustResolveAll[example.SomeStruct]() // -> should panic
}
