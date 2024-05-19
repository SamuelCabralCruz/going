package main

import (
	"errors"
	"github.com/SamuelCabralCruz/went/botox"
	pkg2 "github.com/SamuelCabralCruz/went/botox/internal/example/pkg"
	"github.com/SamuelCabralCruz/went/fn"
)

func main() {
	botox.Register(pkg2.ConfigProvider1)
	botox.Register(pkg2.ConfigProvider2)
	botox.Register(pkg2.ConfigProvider3)
	botox.RegisterSingleton(pkg2.ConfigProvider4)
	botox.RegisterInstance[pkg2.Test1]("coucou")
	botox.RegisterInstance[pkg2.Test2]("caliss")
	botox.RegisterInstance(pkg2.SomeImpl1{})
	botox.RegisterInstance(pkg2.SomeImpl2{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl1{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl2{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl3{})
	botox.RegisterInstance(errors.New("some error"))
	botox.Register(fn.AsResultable(pkg2.NewNested))
	pkg2.Execute()
}
