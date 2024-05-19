package main

import (
	"errors"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg"
	"github.com/samber/mo"
)

func AsResultable[T any](f func() T) func() mo.Result[T] {
	return func() mo.Result[T] { return mo.Ok(f()) }
}

func main() {
	botox.Register(pkg.ConfigProvider1)
	botox.Register(pkg.ConfigProvider2)
	botox.Register(pkg.ConfigProvider3)
	botox.RegisterSingleton(pkg.ConfigProvider4)
	botox.RegisterInstance[pkg.Test1]("coucou")
	botox.RegisterInstance[pkg.Test2]("caliss")
	botox.RegisterInstance(pkg.SomeImpl1{})
	botox.RegisterInstance(pkg.SomeImpl2{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl1{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl2{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl3{})
	botox.RegisterInstance(errors.New("some error"))
	botox.Register(AsResultable(pkg.NewNested))
	pkg.Execute()
}
