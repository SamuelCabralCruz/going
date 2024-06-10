//go:build test

package fn_test

import (
	"errors"
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/tuple/validation"
	"github.com/SamuelCabralCruz/went/fn/typing"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(fn.ValueToSupplier[any], func() {
	var input any
	var observed typing.Supplier[any]

	act := func() {
		observed = fn.ValueToSupplier(input)
	}

	BeforeEach(func() {
		input = "some value"
	})

	It("should return supplier of inputted value", func() {
		act()

		Expect(observed()).To(Equal(input))
	})
})

var _ = DescribeFunction(fn.SupplierToProducer[any], func() {
	var suppliedValue any
	var input typing.Supplier[any]
	var observed typing.Producer[any]

	act := func() {
		observed = fn.SupplierToProducer(input)
	}

	BeforeEach(func() {
		suppliedValue = "some value"
		input = func() any {
			return suppliedValue
		}
	})

	It("should return producer of supplied value and no error", func() {
		act()

		observedValue, observedError := observed()
		Expect(observedValue).To(Equal(suppliedValue))
		Expect(observedError).To(BeNil())
	})
})

var _ = DescribeFunction(fn.ProducerToSupplier[any], func() {
	var producedValue any
	var producedError error
	var input typing.Producer[any]
	var observed typing.Supplier[any]

	act := func() {
		observed = fn.ProducerToSupplier(input)
	}

	BeforeEach(func() {
		producedValue = nil
		producedError = nil
		input = func() (any, error) {
			return producedValue, producedError
		}
	})

	Context("with error returning producer", func() {
		BeforeEach(func() {
			producedError = errors.New("something went wrong")
		})

		It("should panic", func() {
			act()

			Expect(func() { observed() }).To(PanicWith(producedError))
		})
	})

	Context("with value returning producer", func() {
		BeforeEach(func() {
			producedValue = "some value"
		})

		It("should supply produced value", func() {
			act()

			Expect(observed()).To(Equal(producedValue))
		})
	})
})

var _ = DescribeFunction(fn.SupplierToMapper[any], func() {
	var suppliedValue any
	var input typing.Supplier[any]
	var observed typing.Mapper[any]

	act := func() {
		observed = fn.SupplierToMapper(input)
	}

	BeforeEach(func() {
		suppliedValue = "some value"
		input = func() any {
			return suppliedValue
		}
	})

	It("should return mapped of supplied value", func() {
		act()

		Expect(observed("any input")).To(Equal(suppliedValue))
	})
})

var _ = DescribeFunction(fn.MapperToSupplier[any], func() {
	var inputMapper typing.Mapper[any]
	var inputValue any
	var mappedValue any
	var observed typing.Supplier[any]

	act := func() {
		observed = fn.MapperToSupplier(inputMapper, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some input value"
		mappedValue = "some mapped value"
		inputMapper = func(_ any) any {
			return mappedValue
		}
	})

	It("should return supplier of mapped value", func() {
		act()

		Expect(observed()).To(Equal(mappedValue))
	})
})

var _ = DescribeFunction(fn.MapperToTransformer[any], func() {
	var input typing.Mapper[any]
	var inputValue any
	var mappedValue any
	var observed typing.Transformer[any, any]

	act := func() {
		observed = fn.MapperToTransformer(input)
	}

	BeforeEach(func() {
		inputValue = "some input value"
		mappedValue = "some mapped value"
		input = func(_ any) any {
			return mappedValue
		}
	})

	It("should return transformer of mapped value", func() {
		act()

		Expect(observed(inputValue)).To(Equal(mappedValue))
	})
})

var _ = DescribeFunction(fn.AsserterToValidator[any], func() {
	var input typing.Asserter[any]
	var assertedValue any
	var assertedError error
	var observed typing.Validator[any]

	act := func() {
		observed = fn.AsserterToValidator(input)
	}

	BeforeEach(func() {
		assertedValue = nil
		assertedError = nil
		input = func(_ any) (any, error) {
			return assertedValue, assertedError
		}
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			assertedValue = "some value"
		})

		It("should return ok validator of asserted value", func() {
			act()

			validatedValue, validatedOk := observed("some value to validate")
			Expect(validatedOk).To(BeTrue())
			Expect(validatedValue).To(Equal("some value"))
		})
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			assertedError = errors.New("something went wrong")
		})

		It("should return non ok validator of asserted value", func() {
			act()

			validatedValue, validatedOk := observed("some value to validate")
			Expect(validatedOk).To(BeFalse())
			Expect(validatedValue).To(BeZero())
		})
	})
})

var _ = DescribeFunction(fn.ValidatorToPredicate[any], func() {
	var input typing.Validator[any]
	var validatedValue any
	var validatedOk bool
	var observed typing.Predicate[any]

	act := func() {
		observed = fn.ValidatorToPredicate(input)
	}

	BeforeEach(func() {
		validatedValue = nil
		validatedOk = false
		input = func(_ any) (any, bool) {
			return validatedValue, validatedOk
		}
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			validatedValue = "some value"
			validatedOk = true
		})

		It("should return truthy predicate", func() {
			act()

			Expect(observed("some value to predicate")).To(BeTrue())
		})
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			validatedOk = false
		})

		It("should return falsy predicate", func() {
			act()

			Expect(observed("some value to predicate")).To(BeFalse())
		})
	})
})

var _ = DescribeFunction(fn.PredicateToValidator[any], func() {
	var inputtedValue any
	var predicatedValue bool
	var input typing.Predicate[string]
	var observed typing.Validator[string]

	act := func() {
		observed = fn.PredicateToValidator[string](input)
	}

	BeforeEach(func() {
		predicatedValue = false
		input = func(_ string) bool {
			return predicatedValue
		}
	})

	Context("with unexpected type inputted value", func() {
		BeforeEach(func() {
			inputtedValue = 1234
		})

		It("should return non ok validator", func() {
			act()

			validatedValue, validatedOk := observed(inputtedValue)
			Expect(validatedOk).To(BeFalse())
			Expect(validatedValue).To(BeZero())
		})
	})

	Context("with expected type inputted value", func() {
		BeforeEach(func() {
			inputtedValue = "i am a string"
		})

		Context("with truthy predicate", func() {
			BeforeEach(func() {
				predicatedValue = true
			})

			It("should return ok validator", func() {
				act()

				validatedValue, validatedOk := observed(inputtedValue)
				Expect(validatedOk).To(BeTrue())
				Expect(validatedValue).To(Equal(inputtedValue))
			})
		})

		Context("with falsy predicate", func() {
			BeforeEach(func() {
				predicatedValue = false
			})

			It("should return non ok validator", func() {
				act()

				validatedValue, validatedOk := observed(inputtedValue)
				Expect(validatedOk).To(BeFalse())
				Expect(validatedValue).To(BeZero())
			})
		})
	})
})

var _ = DescribeFunction(fn.ValidatorToAsserter[any], func() {
	var validatedValue any
	var validatedOk bool
	var input typing.Validator[any]
	var observed typing.Asserter[any]

	act := func() {
		observed = fn.ValidatorToAsserter(input)
	}

	BeforeEach(func() {
		validatedValue = nil
		validatedOk = false
		input = func(_ any) (any, bool) {
			return validatedValue, validatedOk
		}
	})

	Context("with ok validation", func() {
		BeforeEach(func() {
			validatedOk = true
			validatedValue = "some value"
		})

		It("should return ok asserter of validated value", func() {
			act()

			assertedValue, assertedError := observed("some input value")
			Expect(assertedValue).To(Equal(validatedValue))
			Expect(assertedError).To(BeNil())
		})
	})

	Context("with non ok validation", func() {
		BeforeEach(func() {
			validatedOk = false
		})

		It("should return error asserter", func() {
			act()

			assertedValue, assertedError := observed("some input value")
			Expect(assertedValue).To(BeZero())
			Expect(assertedError).To(BeAssignableToTypeOf(validation.InvalidValueError{}))
		})
	})
})
