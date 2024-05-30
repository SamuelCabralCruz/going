package botox

import (
	"github.com/SamuelCabralCruz/went/botox/internal/it"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/SamuelCabralCruz/went/fn/result"
	"github.com/SamuelCabralCruz/went/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/went/fn/typing"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
)

var container = map[string][]any{}

func RegisterProducer[T any](produce typing.Producer[T]) {
	register(produce, it.Register[T])
}

func RegisterSupplier[T any](supply typing.Supplier[T]) {
	RegisterProducer(fn.SupplierToProducer(supply))
}

func RegisterInstance[T any](instance T) {
	RegisterSupplier(fn.ValueToSupplier(instance))
}

func RegisterSingletonProducer[T any](produce typing.Producer[T]) {
	singletonProducer := func() (*T, error) {
		value, err := produce()
		return &value, err
	}
	register(singletonProducer, it.RegisterSingleton[*T])
}

func RegisterSingletonSupplier[T any](supply typing.Supplier[T]) {
	RegisterSingletonProducer(fn.SupplierToProducer(supply))
}

func RegisterSingletonInstance[T any](instance T) {
	RegisterSingletonSupplier(fn.ValueToSupplier(instance))
}

func register[T any](provider typing.Producer[T], tokenGenerator func(typing.Producer[T]) it.InjectionToken[T]) {
	t := phi.UniqueIdentifier[T]()
	tokens := optional.OfNullable(container[t]).OrElse([]any{})
	container[t] = append(tokens, tokenGenerator(provider))
}

func ResolveAll[T any]() ([]T, error) {
	instances, err := result.Combine(lo.Map(
		container[phi.UniqueIdentifier[T]()],
		func(token any, _ int) result.Result[T] {
			r := token.(it.InjectionToken[T])
			return result.FromAssertion(r.Resolve())
		})...).Get()

	if err != nil {
		return assertion.FromError[[]T](err)
	}

	if len(instances) == 0 {
		return assertion.FromError[[]T](newNoCandidateFoundError(phi.Type[T]()))
	}

	return assertion.FromValue(instances)
}

func MustResolveAll[T any]() []T {
	instances, err := ResolveAll[T]()
	if err != nil {
		panic(err)
	}
	return instances
}

func Resolve[T any]() (T, error) {
	instances, err := ResolveAll[T]()
	if err != nil {
		return assertion.FromError[T](err)
	}
	if len(instances) > 1 {
		return assertion.FromError[T](newTooManyCandidatesFoundError(phi.Type[T](), len(instances)))
	}
	return assertion.FromValue(instances[0])
}

func MustResolve[T any]() T {
	instance, err := Resolve[T]()
	if err != nil {
		panic(err)
	}
	return instance
}

func Clear() {
	container = map[string][]any{}
}
