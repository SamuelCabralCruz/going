package container

import (
	internal "github.com/SamuelCabralCruz/going/botox/internal/container"
	"github.com/SamuelCabralCruz/going/botox/internal/it"
	"github.com/SamuelCabralCruz/going/fn"
	"github.com/SamuelCabralCruz/going/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/going/fn/typing"
	"github.com/SamuelCabralCruz/going/phi"
)

func New() *Container {
	return &Container{ref: internal.NewContainer()}
}

type Container struct {
	ref *internal.Container
}

func RegisterProducer[T any](container *Container, produce typing.Producer[T]) {
	register(container, produce, it.Register[T])
}

func RegisterSupplier[T any](container *Container, supply typing.Supplier[T]) {
	RegisterProducer(container, fn.SupplierToProducer(supply))
}

func RegisterInstance[T any](container *Container, instance T) {
	RegisterSupplier(container, fn.ValueToSupplier(instance))
}

func RegisterSingletonProducer[T any](
	container *Container,
	produce typing.Producer[T],
) {
	singletonProducer := func() (*T, error) {
		value, err := produce()
		return &value, err
	}
	register(container, singletonProducer, it.RegisterSingleton[*T])
}

func RegisterSingletonSupplier[T any](container *Container, supply typing.Supplier[T]) {
	RegisterSingletonProducer(container, fn.SupplierToProducer(supply))
}

func RegisterSingletonInstance[T any](container *Container, instance T) {
	RegisterSingletonSupplier(container, fn.ValueToSupplier(instance))
}

func register[T any](
	container *Container,
	provider typing.Producer[T],
	tokenGenerator func(typing.Producer[T]) it.InjectionToken[T],
) {
	token := tokenGenerator(provider)
	internal.RegisterToken(container.ref, token)
}

func ResolveAll[T any](container *Container) ([]T, error) {
	instances, err := internal.ResolveTokens[T](container.ref)
	if err != nil {
		return assertion.FromError[[]T](err)
	}
	if len(instances) == 0 {
		return assertion.FromError[[]T](newNoCandidateFoundError(phi.Type[T]()))
	}
	return assertion.FromValue(instances)
}

func MustResolveAll[T any](container *Container) []T {
	return assertion.GetOrPanic(ResolveAll[T](container))
}

func Resolve[T any](container *Container) (T, error) {
	instances, err := ResolveAll[T](container)
	if err != nil {
		return assertion.FromError[T](err)
	}
	if len(instances) > 1 {
		return assertion.FromError[T](newTooManyCandidatesFoundError(phi.Type[T](), len(instances)))
	}
	return assertion.FromValue(instances[0])
}

func MustResolve[T any](container *Container) T {
	return assertion.GetOrPanic(Resolve[T](container))
}

func Unregister[T any](container *Container) {
	internal.Unregister[T](container.ref)
}

func Reset(container *Container) {
	internal.Reset(container.ref)
}

func Clone(container *Container) *Container {
	return &Container{ref: internal.Clone(container.ref)}
}
