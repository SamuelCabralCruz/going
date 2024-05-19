package pkg

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model1"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model2"
	"github.com/samber/lo"
)

func Execute() {
	lo.ForEach(botox.MustResolveAll[model1.Config](), func(c model1.Config, _ int) {
		fmt.Println(fmt.Sprintf("%+v", c))
	})
	lo.ForEach(botox.MustResolveAll[model1.Config](), func(c model1.Config, _ int) {
		fmt.Println(fmt.Sprintf("%+v", c))
	})

	fmt.Println(fmt.Sprintf("%+v", botox.MustResolve[model2.Config]()))
	fmt.Println(fmt.Sprintf("%+v", botox.MustResolve[Test1]()))
	fmt.Println(fmt.Sprintf("%+v", botox.MustResolve[Test2]()))

	fmt.Println("interfaces")
	lo.ForEach(botox.MustResolveAll[SomeInterface](), func(c SomeInterface, _ int) {
		fmt.Println(fmt.Sprintf("%+v", c.Coucou()))
	})

	fmt.Println(botox.MustResolve[error]().Error())

	fmt.Println(botox.MustResolve[Nested]().ToString())
}
