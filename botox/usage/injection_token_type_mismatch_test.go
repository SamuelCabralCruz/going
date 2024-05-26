//go:build test

package usage_test

// TODO: test registration of ptr singleton vs non singleton

// TODO: test mutation of singleton vs non singleton

// TODO: test producer returning error

// TODO: test panic during supplier
// TODO: test panic during supplier singleton
// TODO: test panic during producer
// TODO: test panic during producer singleton

// TODO: impossible...
// TODO: Remove unnecessary logic associated to this use case
//var _ = DescribeFunction(botox.RegisterInstance[any], func() {
//	//var instance1 pkg1.SomeStruct
//	var instance2 pkg2.SomeStruct
//
//	register := func() {
//		botox.RegisterInstance[any](instance2)
//	}
//
//	BeforeEach(func() {
//		//instance1 = pkg1.SomeStruct{}
//		instance2 = pkg2.SomeStruct{}
//	})
//
//	AfterEach(func() {
//		botox.Clear()
//	})
//
//	DescribeFunction(botox.MustResolve[any], func() {
//		resolve := func() {
//			botox.MustResolve[pkg1.SomeStruct]()
//		}
//
//		It("should panic", func() {
//			register()
//
//			Expect(resolve).To(PanicWith(BeAssignableToTypeOf(botox.NoCandidateFoundError{})))
//		})
//	})
//})
