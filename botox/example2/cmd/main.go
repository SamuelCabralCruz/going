package main

import (
	"errors"
	"github.com/SamuelCabralCruz/went/botox"
	pkg2 "github.com/SamuelCabralCruz/went/botox/example2/pkg"
	"github.com/SamuelCabralCruz/went/botox/example2/pkg/model1"
	"github.com/SamuelCabralCruz/went/botox/example2/pkg/model2"
	"github.com/SamuelCabralCruz/went/fn"
	"time"
)

func main() {
	botox.RegisterInstance(model1.Config{
		A: "register instance",
		B: time.Now().Nanosecond(),
	})
	botox.RegisterSupplier(func() model1.Config {
		return model1.Config{
			A: "register supplier",
			B: time.Now().Nanosecond(),
		}
	})
	botox.RegisterProducer(func() (model1.Config, error) {
		return model1.Config{
			A: "register producer",
			B: time.Now().Nanosecond(),
		}, nil
	})
	botox.RegisterSingletonSupplier(func() model2.Config {
		return model2.Config{
			A: "register singleton supplier",
			B: time.Now().Nanosecond(),
		}
	})
	botox.RegisterSingletonProducer(func() (model2.Config, error) {
		return model2.Config{
			A: "register singleton producer",
			B: time.Now().Nanosecond(),
		}, nil
	})

	botox.RegisterInstance[pkg2.Test1]("test1")
	botox.RegisterInstance[pkg2.Test2]("test2")

	botox.RegisterInstance(pkg2.SomeImpl1{})
	botox.RegisterInstance(pkg2.SomeImpl2{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl1{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl2{})
	botox.RegisterInstance[pkg2.SomeInterface](pkg2.SomeImpl3{})

	botox.RegisterInstance(errors.New("some error"))
	botox.RegisterProducer(fn.ToProducer(pkg2.NewNested))

	botox.RegisterInstance(pkg2.Mutate{})                                        // T
	botox.RegisterSingletonInstance(pkg2.Mutate{})                               // *T
	botox.RegisterSupplier(func() pkg2.Mutate { return pkg2.Mutate{} })          // T
	botox.RegisterSingletonSupplier(func() pkg2.Mutate { return pkg2.Mutate{} }) // *T

	botox.RegisterInstance(&pkg2.Mutate{})                                         // *T
	botox.RegisterSingletonInstance(&pkg2.Mutate{})                                // **T
	botox.RegisterSupplier(func() *pkg2.Mutate { return &pkg2.Mutate{} })          // *T
	botox.RegisterSingletonSupplier(func() *pkg2.Mutate { return &pkg2.Mutate{} }) // **T

	pkg2.Execute()
}
