package it

import (
	"github.com/samber/mo"
)

type Provider[T any] func() mo.Result[T]

type InjectionToken[T any] interface {
	Resolve() mo.Result[T]
	MustResolve() T
}
