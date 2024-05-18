package botox

import (
	"github.com/SamuelCabralCruz/went/botox/internal/it"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/phi"
	"github.com/samber/lo"
	"github.com/samber/mo"
)

var container = map[string][]any{}

func RegisterInstance[T any](instance T) {
	Register(func() mo.Result[T] { return mo.Ok(instance) })
}

func Register[T any](provider it.Provider[T]) {
	register(provider, it.Register[T])
}

func RegisterSingleton[T any](provider it.Provider[T]) {
	register(provider, it.RegisterSingleton[T])
}

func register[T any](provider it.Provider[T], tokenGenerator func(it.Provider[T]) it.InjectionToken[T]) {
	t := phi.UniqueIdentifier[T]()
	tokens := mo.EmptyableToOption(container[t]).OrElse([]any{})
	container[t] = append(tokens, tokenGenerator(provider))
}

func ResolveAll[T any]() (instances []T, err error) {
	instances, err = fn.Accumulate(lo.Map(
		container[phi.UniqueIdentifier[T]()],
		func(token any, _ int) mo.Result[T] {
			if r, ok := token.(it.InjectionToken[T]); ok {
				return r.Resolve()
			}
			return mo.Err[T](newProvidingLoopError())
		})...).Get()

	if err != nil {
		return fn.ErrorHasTuple[[]T](err)
	}

	if len(instances) == 0 {
		return fn.ErrorHasTuple[[]T](newNoCandidateFoundError(phi.Type[T]()))
	}

	return fn.ValueHasTuple(instances)
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
		return fn.ErrorHasTuple[T](err)
	}
	if len(instances) > 1 {
		return fn.ErrorHasTuple[T](newTooManyCandidatesFoundError(phi.Type[T](), len(instances)))
	}
	return fn.ValueHasTuple(instances[0])
}

func MustResolve[T any]() T {
	instance, err := Resolve[T]()
	if err != nil {
		panic(err)
	}
	return instance
}
