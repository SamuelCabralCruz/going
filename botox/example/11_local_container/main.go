//go:build example

package main

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/container"
	"github.com/SamuelCabralCruz/went/botox/example/00_boilerplate"
)

func main() {
	// Until now, we only covered the usage of a global container.
	// However, Botox also allows the creation of local container.

	// This feature can be particularly useful when you desire to create multiple
	// injection contexts that will be independent of one another.

	// There are multiple ways of creating a local container.
	// 1. Out of thin air 💨
	//    `container.New()`
	// 2. Localizing the global container 🌎
	//    `botox.Localize()`
	// 3. Creating a child of local container 👶
	//    `container.Clone(someLocalContainer)`

	// Of course, all the features showcased for the global container are also
	// available for the local container.

	// REGISTRATION
	botox.RegisterSingletonSupplier(func() example.SomeStruct {
		fmt.Println("singleton supplier invoked")
		return example.SomeStruct{}
	})

	localContainer := container.New() // -> create virgin local container
	container.RegisterInstance[example.SomeImplementation2](localContainer, example.SomeImplementation2{})

	localizedGlobalContainer := botox.Localize() // -> create local container from global
	container.RegisterSupplier[example.SomeImplementation1](localizedGlobalContainer, func() example.SomeImplementation1 {
		return example.SomeImplementation1{}
	})
	// RESOLUTION
	_, err := botox.Resolve[example.SomeImplementation1]() // should return an error
	fmt.Println(err.Error())

	_, err = container.Resolve[example.SomeImplementation1](localContainer) // should return an error
	fmt.Println(err.Error())
	container.MustResolve[example.SomeImplementation2](localContainer) // should not panic

	container.MustResolve[example.SomeImplementation1](localizedGlobalContainer) // should not panic

	// Side Note:
	// The singletons' resolution state of a local container will be inherited from
	// the parent container.
	// This means that if a singleton has already been resolved by the parent container,
	// the same resolution result will be returned.
	// Similarly, if a singleton was not yet resolved, the provider will be invoked.
	botox.MustResolve[*example.SomeStruct]()                             // -> should execute the provider
	botox.MustResolve[*example.SomeStruct]()                             // -> should not execute the provider
	container.MustResolve[*example.SomeStruct](localizedGlobalContainer) // -> should execute the provider
	container.MustResolve[*example.SomeStruct](localizedGlobalContainer) // -> should not execute the provider
	afterResolution := botox.Localize()
	container.MustResolve[*example.SomeStruct](afterResolution) // -> should not execute the provider
}
