package main

import (
	"errors"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model1"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model2"
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

	botox.RegisterInstance[pkg.Test1]("test1")
	botox.RegisterInstance[pkg.Test2]("test2")

	botox.RegisterInstance(pkg.SomeImpl1{})
	botox.RegisterInstance(pkg.SomeImpl2{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl1{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl2{})
	botox.RegisterInstance[pkg.SomeInterface](pkg.SomeImpl3{})

	botox.RegisterInstance(errors.New("some error"))
	botox.RegisterProducer(fn.ToProducer(pkg.NewNested))

	botox.RegisterInstance(pkg.Mutate{})                                       // T
	botox.RegisterSingletonInstance(pkg.Mutate{})                              // *T
	botox.RegisterSupplier(func() pkg.Mutate { return pkg.Mutate{} })          // T
	botox.RegisterSingletonSupplier(func() pkg.Mutate { return pkg.Mutate{} }) // *T

	botox.RegisterInstance(&pkg.Mutate{})                                        // *T
	botox.RegisterSingletonInstance(&pkg.Mutate{})                               // **T
	botox.RegisterSupplier(func() *pkg.Mutate { return &pkg.Mutate{} })          // *T
	botox.RegisterSingletonSupplier(func() *pkg.Mutate { return &pkg.Mutate{} }) // **T

	pkg.Execute()
}
