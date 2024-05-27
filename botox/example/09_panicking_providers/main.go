//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// Since we have trust issues, we will wrap your dependency providers to
	// recover from any potential panic!

	// REGISTRATION
	botox.RegisterSupplier(func() example.SomeImplementation1 {
		panic("something went wrong - supplier")
	})
	botox.RegisterProducer(func() (example.SomeImplementation1, error) {
		panic("something went wrong - producer")
	})
	// RESOLUTION
	instances, err := botox.ResolveAll[example.SomeImplementation1]()
	fmt.Println(fmt.Sprintf("nb instances: %d", len(instances)))
	fmt.Println(err.Error())
}
