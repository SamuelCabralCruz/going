package botox

import (
	"github.com/SamuelCabralCruz/went/botox/internal/it"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/SamuelCabralCruz/went/fn/result"
	"github.com/SamuelCabralCruz/went/fn/tuple"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
)

var container = map[string][]any{}

func RegisterProducer[T any](produce fn.Producer[T]) {
	register(produce, it.Register[T])
}

func RegisterSupplier[T any](supply fn.Supplier[T]) {
	RegisterProducer(fn.ToProducer(supply))
}

func RegisterInstance[T any](instance T) {
	RegisterSupplier(fn.ToSupplier(instance))
}

func RegisterSingletonProducer[T any](produce fn.Producer[T]) {
	singletonProducer := func() (*T, error) {
		value, err := produce()
		return &value, err
	}
	register(singletonProducer, it.RegisterSingleton[*T])
}

func RegisterSingletonSupplier[T any](supply fn.Supplier[T]) {
	RegisterSingletonProducer(fn.ToProducer(supply))
}

func RegisterSingletonInstance[T any](instance T) {
	RegisterSingletonSupplier(fn.ToSupplier(instance))
}

func register[T any](provider fn.Producer[T], tokenGenerator func(fn.Producer[T]) it.InjectionToken[T]) {
	t := phi.UniqueIdentifier[T]()
	tokens := optional.OfNullable(container[t]).OrElse([]any{})
	container[t] = append(tokens, tokenGenerator(fn.ToSafeProducer(provider)))
}

func ResolveAll[T any]() ([]T, error) {
	instances, err := result.Combine(lo.Map(
		container[phi.UniqueIdentifier[T]()],
		func(token any, _ int) result.Result[T] {
			r := token.(it.InjectionToken[T])
			return result.FromTuple(r.Resolve())
		})...).Get()

	if err != nil {
		return tuple.FromError[[]T](err)
	}

	if len(instances) == 0 {
		return tuple.FromError[[]T](newNoCandidateFoundError(phi.Type[T]()))
	}

	return tuple.FromValue(instances)
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
		return tuple.FromError[T](err)
	}
	if len(instances) > 1 {
		return tuple.FromError[T](newTooManyCandidatesFoundError(phi.Type[T](), len(instances)))
	}
	return tuple.FromValue(instances[0])
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
