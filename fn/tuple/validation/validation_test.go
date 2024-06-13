//go:build test

package validation_test

import (
	"errors"
	"github.com/SamuelCabralCruz/went/fn/tuple/validation"
	"github.com/SamuelCabralCruz/went/fn/typing"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(validation.IgnoreOk[any], func() {
	var inputValue any
	var inputOk bool
	var observed any

	act := func() {
		observed = validation.IgnoreOk(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputOk = true
	})

	It("should return validated value", func() {
		act()

		Expect(observed).To(Equal(inputValue))
	})
})

var _ = DescribeFunction(validation.IgnoreValue[any], func() {
	var inputValue any
	var inputOk bool
	var observed any

	act := func() {
		observed = validation.IgnoreValue(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputOk = true
	})

	It("should return validated ok", func() {
		act()

		Expect(observed).To(Equal(inputOk))
	})
})

var _ = DescribeFunction(validation.GetOrEmpty[any], func() {
	var inputValue any
	var inputOk bool
	var observed any

	act := func() {
		observed = validation.GetOrEmpty(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = nil
		inputOk = false
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
		})

		It("should return empty value", func() {
			act()

			Expect(observed).To(BeZero())
		})
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should return validated value", func() {
			act()

			Expect(observed).To(Equal(inputValue))
		})
	})
})

var _ = DescribeFunction(validation.GetOrPanic[any], func() {
	var inputValue any
	var inputOk bool
	var observed any

	act := func() {
		observed = validation.GetOrPanic(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = nil
		inputOk = false
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
		})

		It("should panic", func() {
			Expect(act).To(PanicWith(BeAssignableToTypeOf(validation.InvalidValueError{})))
		})
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should return validated value", func() {
			act()

			Expect(observed).To(Equal(inputValue))
		})
	})
})

var _ = DescribeFunction(validation.PanicIfNotOk[any], func() {
	var inputValue any
	var inputOk bool

	act := func() {
		validation.PanicIfNotOk(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = nil
		inputOk = false
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
		})

		It("should panic", func() {
			Expect(act).To(PanicWith(BeAssignableToTypeOf(validation.InvalidValueError{})))
		})
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should not panic", func() {
			Expect(act).NotTo(Panic())
		})
	})
})

var _ = DescribeFunction(validation.Ok[any], func() {
	var input any
	var observedValue any
	var observedOk bool

	act := func() {
		observedValue, observedOk = validation.Ok(input)
	}

	BeforeEach(func() {
		input = "some value"
	})

	It("should return ok validation", func() {
		act()

		Expect(observedValue).To(Equal(input))
		Expect(observedOk).To(BeTrue())
	})
})

var _ = DescribeFunction(validation.NotOk[any], func() {
	var observedValue any
	var observedOk bool

	act := func() {
		observedValue, observedOk = validation.NotOk[any]()
	}

	It("should return non ok validation", func() {
		act()

		Expect(observedValue).To(BeZero())
		Expect(observedOk).To(BeFalse())
	})
})

var _ = DescribeFunction(validation.FromReversed[any], func() {
	var inputValue any
	var inputOk bool
	var observedValue any
	var observedOk bool

	act := func() {
		observedValue, observedOk = validation.FromReversed[any](inputOk, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputOk = true
	})

	It("should return corresponding validation", func() {
		act()

		Expect(observedValue).To(Equal(inputValue))
		Expect(observedOk).To(Equal(inputOk))
	})
})

var _ = DescribeFunction(validation.ToAssertion[any], func() {
	var inputValue any
	var inputOk bool
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = validation.ToAssertion(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = nil
		inputOk = false
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should return ok assertion", func() {
			act()

			Expect(observedValue).To(Equal(inputValue))
			Expect(observedError).To(BeNil())
		})
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
			inputValue = "some ignored value"
		})

		It("should return error assertion", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(BeAssignableToTypeOf(validation.InvalidValueError{}))
		})
	})
})

var _ = DescribeFunction(validation.ToAssertionWithError[any], func() {
	var inputValue any
	var inputOk bool
	var inputError error
	var observed func(error) (any, error)

	act := func() {
		observed = validation.ToAssertionWithError(inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = nil
		inputOk = false
		inputError = errors.New("something went wrong")
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should return ok assertion", func() {
			act()

			observedValue, observedError := observed(inputError)
			Expect(observedValue).To(Equal(inputValue))
			Expect(observedError).To(BeNil())
		})
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
			inputValue = "some ignored value"
		})

		It("should return error assertion", func() {
			act()

			observedValue, observedError := observed(inputError)
			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(inputError))
		})
	})
})

var _ = DescribeFunction(validation.Switch[any, any], func() {
	var inputValue string
	var inputOk bool
	var receivedOnOk string
	var transformedOnOk int
	var suppliedOnNotOk int
	var onOk typing.Transformer[string, int]
	var onNotOk typing.Supplier[int]
	var observed func(typing.Transformer[string, int], typing.Supplier[int]) int

	act := func() {
		observed = validation.Switch[string, int](inputValue, inputOk)
	}

	BeforeEach(func() {
		inputValue = ""
		inputOk = false
		receivedOnOk = ""
		transformedOnOk = 1234
		suppliedOnNotOk = 4321
		onOk = func(v string) int {
			receivedOnOk = v
			return transformedOnOk
		}
		onNotOk = func() int {
			return suppliedOnNotOk
		}
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			inputOk = true
			inputValue = "some value"
		})

		It("should transform validated value using on ok transformer", func() {
			act()

			Expect(observed(onOk, onNotOk)).To(Equal(transformedOnOk))
			Expect(receivedOnOk).To(Equal(inputValue))
		})
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			inputOk = false
		})

		It("should supply value using on non ok supplier", func() {
			act()

			Expect(observed(onOk, onNotOk)).To(Equal(suppliedOnNotOk))
		})
	})
})
