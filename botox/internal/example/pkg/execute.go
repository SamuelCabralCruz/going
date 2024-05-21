package pkg

import (
	"fmt"
	"github.com/SamuelCabralCruz/went/botox"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model1"
	"github.com/SamuelCabralCruz/went/botox/internal/example/pkg/model2"
	"github.com/SamuelCabralCruz/went/fn"
)

func Execute() {
	// TODO: transform as proper tests
	reportMustResolve(botox.MustResolveAll[model1.Config])
	reportMustResolve(botox.MustResolveAll[model1.Config])
	reportMustResolve(botox.MustResolve[model1.Config])

	reportMustResolve(botox.MustResolveAll[model2.Config])
	reportMustResolve(botox.MustResolveAll[model2.Config])
	reportMustResolve(botox.MustResolve[model2.Config])

	reportMustResolve(botox.MustResolve[Test1])
	reportMustResolve(botox.MustResolve[Test2])
	reportResolve(botox.Resolve[Test1])
	reportResolve(botox.Resolve[Test2])

	reportMustResolve(botox.MustResolveAll[SomeInterface])

	reportMustResolve(botox.MustResolve[error])

	reportMustResolve(botox.MustResolve[Nested])

	mutableValues := botox.MustResolveAll[Mutate]()
	fmt.Println(mutableValues)
	//mutableValues[0].Increment()
	//mutableValues[0].A = "incremented once"
	//mutableValues[1].Increment()
	//mutableValues[1].Increment()
	//mutableValues[1].A = "incremented twice"
	//fmt.Println(botox.MustResolveAll[Mutate]()) // no effect

	mutableRefs := botox.MustResolveAll[*Mutate]()
	fmt.Println(mutableRefs)
	//lo.ForEach(mutableRefs, func(item *Mutate, _ int) {
	//	fmt.Println(*item)
	//})
	//mutableRefs[0].Increment()
	//mutableRefs[0].A = "incremented once"
	//mutableRefs[1].Increment()
	//mutableRefs[1].Increment()
	//mutableRefs[1].A = "incremented twice"
	//mutableRefs[2].Increment()
	//mutableRefs[2].Increment()
	//mutableRefs[2].Increment()
	//mutableRefs[2].A = "incremented thrice"
	//mutableRefs[3].Increment()
	//mutableRefs[3].Increment()
	//mutableRefs[3].Increment()
	//mutableRefs[3].Increment()
	//mutableRefs[3].A = "incremented four times"
	//mutableRefs[4].Increment()
	//mutableRefs[4].Increment()
	//mutableRefs[4].Increment()
	//mutableRefs[4].Increment()
	//mutableRefs[4].Increment()
	//mutableRefs[4].A = "incremented five times"
	//mutableRefsAfter := botox.MustResolveAll[*Mutate]()
	//lo.ForEach(mutableRefsAfter, func(item *Mutate, _ int) {
	//	fmt.Println(*item)
	//})

	mutableRefRefs := botox.MustResolveAll[**Mutate]()
	fmt.Println(mutableRefRefs)
}

func reportMustResolve[T any](resolve fn.Supplier[T]) {
	reportResolve(fn.ToProducer(resolve))
}

func reportResolve[T any](resolve fn.Producer[T]) {
	value, err := fn.Try(resolve)
	fmt.Printf("%+v, %+v\n", value, err)
}
