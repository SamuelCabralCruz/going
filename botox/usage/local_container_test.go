//go:build test

package usage_test

import (
  "github.com/SamuelCabralCruz/went/botox"
  "github.com/SamuelCabralCruz/went/botox/container"
  "github.com/SamuelCabralCruz/went/botox/usage/fixture"
  . "github.com/SamuelCabralCruz/went/kinggo"
  . "github.com/onsi/ginkgo/v2"
  . "github.com/onsi/gomega"
)

var _ = DescribeType[container.Container](func() {
  DescribeFunction(container.New, func() {
    var container1 *container.Container
    var container2 *container.Container
    var container1SingletonInvoked bool
    var container2SingletonInvoked bool

    BeforeEach(func() {
      container1 = container.New()
      container2 = container.New()
      container1SingletonInvoked = false
      container2SingletonInvoked = false

      container.RegisterInstance(container1, fixture.SomeImplementation1{})
      container.RegisterInstance(container1, fixture.SomeImplementation2{})
      container.RegisterSingletonSupplier(container1, func() fixture.Stateful {
        container1SingletonInvoked = true
        return fixture.Stateful{}
      })

      container.RegisterInstance(container2, fixture.SomeImplementation3{})
      container.RegisterSingletonSupplier(container2, func() fixture.Stateful {
        container2SingletonInvoked = true
        return fixture.Stateful{}
      })
    })

    It("should create independent containers", func() {
      Expect(func() { container.MustResolve[fixture.SomeImplementation1](container1) }).NotTo(Panic())
      Expect(func() { container.MustResolve[fixture.SomeImplementation2](container1) }).NotTo(Panic())
      Expect(func() { container.MustResolve[fixture.SomeImplementation3](container1) }).To(Panic())
      Expect(func() { container.MustResolve[fixture.SomeImplementation1](container2) }).To(Panic())
      Expect(func() { container.MustResolve[fixture.SomeImplementation2](container2) }).To(Panic())
      Expect(func() { container.MustResolve[fixture.SomeImplementation3](container2) }).NotTo(Panic())

      Expect(func() { container.MustResolve[*fixture.Stateful](container1) }).NotTo(Panic())
      Expect(container1SingletonInvoked).To(BeTrue())
      Expect(container2SingletonInvoked).To(BeFalse())
    })
  })
})

var _ = DescribeFunction(container.Clone, func() {
  var input *container.Container
  var observed *container.Container

  act := func() {
    observed = container.Clone(input)
  }

  BeforeEach(func() {
    input = container.New()

    container.RegisterInstance(input, fixture.SomeImplementation1{})
    container.RegisterInstance(input, fixture.SomeImplementation2{})
    container.RegisterSingletonSupplier(input, func() fixture.Stateful {
      return fixture.Stateful{}
    })
    container.RegisterSingletonSupplier(input, func() fixture.Stateless {
      return fixture.NewStateless()
    })
  })

  It("should deep copy container state", func() {
    stateful := container.MustResolve[*fixture.Stateful](input)
    stateful.Mutate()
    stateful.Mutate()
    stateful.Mutate()

    act()

    Expect(func() { container.MustResolve[fixture.SomeImplementation1](observed) }).NotTo(Panic())
    Expect(func() { container.MustResolve[fixture.SomeImplementation2](observed) }).NotTo(Panic())
    Expect(func() { container.MustResolve[fixture.SomeImplementation3](observed) }).To(Panic())
    inputStateless := container.MustResolve[*fixture.Stateless](input)
    Expect(container.MustResolve[*fixture.Stateful](observed).Count()).To(Equal(3))
    Expect(container.MustResolve[*fixture.Stateless](observed)).NotTo(Equal(inputStateless))
  })
})

var _ = DescribeFunction(botox.Localize, func() {
  var observed *container.Container

  act := func() {
    observed = botox.Localize()
  }

  BeforeEach(func() {
    botox.RegisterInstance(fixture.SomeImplementation1{})
    botox.RegisterInstance(fixture.SomeImplementation2{})
    botox.RegisterSingletonSupplier(func() fixture.Stateful {
      return fixture.Stateful{}
    })
    botox.RegisterSingletonSupplier(func() fixture.Stateless {
      return fixture.NewStateless()
    })
  })

  AfterEach(func() {
    botox.Reset()
  })

  It("should deep copy global container state", func() {
    stateful := botox.MustResolve[*fixture.Stateful]()
    stateful.Mutate()
    stateful.Mutate()
    stateful.Mutate()

    act()

    Expect(func() { container.MustResolve[fixture.SomeImplementation1](observed) }).NotTo(Panic())
    Expect(func() { container.MustResolve[fixture.SomeImplementation2](observed) }).NotTo(Panic())
    Expect(func() { container.MustResolve[fixture.SomeImplementation3](observed) }).To(Panic())
    inputStateless := botox.MustResolve[*fixture.Stateless]()
    Expect(container.MustResolve[*fixture.Stateful](observed).Count()).To(Equal(3))
    Expect(container.MustResolve[*fixture.Stateless](observed)).NotTo(Equal(inputStateless))
  })
})
