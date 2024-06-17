package container

import (
	"github.com/SamuelCabralCruz/going/botox/internal/it"
	"github.com/SamuelCabralCruz/going/fn/optional"
	"github.com/SamuelCabralCruz/going/fn/result"
	"github.com/SamuelCabralCruz/going/phi"
	"github.com/samber/lo"
	"reflect"
)

func NewContainer() *Container {
	return &Container{
		registrations: map[string][]any{},
	}
}

type Container struct {
	registrations map[string][]any
}

func RegisterToken[T any](container *Container, token it.InjectionToken[T]) {
	t := phi.UniqueIdentifier[T]()
	tokens := optional.OfNullable(container.registrations[t]).OrElse([]any{})
	container.registrations[t] = append(tokens, token)
}

func ResolveTokens[T any](container *Container) ([]T, error) {
	tokens := container.registrations[phi.UniqueIdentifier[T]()]
	resolveToken := func(token any, _ int) result.Result[T] {
		r := token.(it.InjectionToken[T])
		return result.FromAssertion(r.Resolve())
	}
	resolved := lo.Map(tokens, resolveToken)
	return result.Combine(resolved...).Get()
}

func Unregister[T any](container *Container) {
	container.registrations[phi.UniqueIdentifier[T]()] = []any{}
}

func Reset(container *Container) {
	container.registrations = map[string][]any{}
}

func Clone(container *Container) *Container {
	cloned := map[string][]any{}
	for k, v := range container.registrations {
		cloned[k] = lo.Map(v, func(item any, _ int) any {
			cloneFunctionName := phi.FunctionName(it.InjectionToken[any].Clone)
			return reflect.
				ValueOf(item).
				MethodByName(cloneFunctionName).
				Call(nil)[0].
				Interface()
		})
	}
	return &Container{registrations: cloned}
}
