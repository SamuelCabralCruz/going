package botox

import (
	"github.com/SamuelCabralCruz/went/botox/container"
	"github.com/SamuelCabralCruz/went/fn/typing"
)

var globalContainer = container.New()

func RegisterProducer[T any](produce typing.Producer[T]) {
	container.RegisterProducer(globalContainer, produce)
}

func RegisterSupplier[T any](supply typing.Supplier[T]) {
	container.RegisterSupplier(globalContainer, supply)
}

func RegisterInstance[T any](instance T) {
	container.RegisterInstance(globalContainer, instance)
}

func RegisterSingletonProducer[T any](produce typing.Producer[T]) {
	container.RegisterSingletonProducer(globalContainer, produce)
}

func RegisterSingletonSupplier[T any](supply typing.Supplier[T]) {
	container.RegisterSingletonSupplier(globalContainer, supply)
}

func RegisterSingletonInstance[T any](instance T) {
	container.RegisterSingletonInstance(globalContainer, instance)
}

func ResolveAll[T any]() ([]T, error) {
	return container.ResolveAll[T](globalContainer)
}

func MustResolveAll[T any]() []T {
	return container.MustResolveAll[T](globalContainer)
}

func Resolve[T any]() (T, error) {
	return container.Resolve[T](globalContainer)
}

func MustResolve[T any]() T {
	return container.MustResolve[T](globalContainer)
}

func Unregister[T any]() {
	container.Unregister[T](globalContainer)
}

func Reset() {
	container.Reset(globalContainer)
}

func Localize() *container.Container {
	return container.Clone(globalContainer)
}
