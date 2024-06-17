//go:build test

package fn_test

import (
	"errors"
	"github.com/SamuelCabralCruz/going/fn"
	"github.com/SamuelCabralCruz/going/fn/typing"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(fn.SafeProducer[any], func() {
	var input typing.Producer[any]
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = fn.SafeProducer(input)
	}

	BeforeEach(func() {
		observedValue = nil
		observedError = nil
	})

	Context("with panicking producer", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			input = func() (any, error) {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking producer", func() {
		Context("with producer returning error", func() {
			returnedError := errors.New("something went wrong")

			BeforeEach(func() {
				input = func() (any, error) {
					return nil, returnedError
				}
			})

			It("should return error", func() {
				act()

				Expect(observedValue).To(BeZero())
				Expect(observedError).To(Equal(returnedError))
			})
		})

		Context("with producer returning value", func() {
			value := "some value"

			BeforeEach(func() {
				input = func() (any, error) {
					return value, nil
				}
			})

			It("should return produced value", func() {
				act()

				Expect(observedValue).To(Equal(value))
				Expect(observedError).To(BeNil())
			})
		})
	})
})

var _ = DescribeFunction(fn.SafeSupplier[any], func() {
	var input typing.Supplier[any]
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = fn.SafeSupplier(input)
	}

	BeforeEach(func() {
		observedValue = nil
		observedError = nil
	})

	Context("with panicking supplier", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			input = func() any {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking supplier", func() {
		suppliedValue := "some supplied value"

		BeforeEach(func() {
			input = func() any {
				return suppliedValue
			}
		})

		It("should return supplied value", func() {
			act()

			Expect(observedValue).To(Equal(suppliedValue))
			Expect(observedError).To(BeNil())
		})
	})
})

var _ = DescribeFunction(fn.SafeMapper[any], func() {
	var inputMapper typing.Mapper[string]
	var inputValue string
	var observedValue string
	var observedError error

	act := func() {
		observedValue, observedError = fn.SafeMapper(inputMapper, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some value"
		observedValue = ""
		observedError = nil
	})

	Context("with panicking mapper", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			inputMapper = func(_ string) string {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking mapper", func() {
		mappedValue := "some mapped value"

		BeforeEach(func() {
			inputMapper = func(_ string) string {
				return mappedValue
			}
		})

		It("should return mapped value", func() {
			act()

			Expect(observedValue).To(Equal(mappedValue))
			Expect(observedError).To(BeNil())
		})
	})
})

var _ = DescribeFunction(fn.SafeTransformer[any, any], func() {
	var inputTransformer typing.Transformer[string, int]
	var inputValue string
	var observedValue int
	var observedError error

	act := func() {
		observedValue, observedError = fn.SafeTransformer(inputTransformer, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some value"
		observedValue = 0
		observedError = nil
	})

	Context("with panicking transformer", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			inputTransformer = func(_ string) int {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking transformer", func() {
		transformedValue := 34

		BeforeEach(func() {
			inputTransformer = func(_ string) int {
				return transformedValue
			}
		})

		It("should return transformed value", func() {
			act()

			Expect(observedValue).To(Equal(transformedValue))
			Expect(observedError).To(BeNil())
		})
	})
})

var _ = DescribeFunction(fn.SafePredicate[any], func() {
	var inputPredicate typing.Predicate[string]
	var inputValue string
	var observedValue bool
	var observedError error

	act := func() {
		observedValue, observedError = fn.SafePredicate(inputPredicate, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some value"
		observedValue = false
		observedError = nil
	})

	Context("with panicking predicate", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			inputPredicate = func(_ string) bool {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking predicate", func() {
		predicatedValue := true

		BeforeEach(func() {
			inputPredicate = func(_ string) bool {
				return predicatedValue
			}
		})

		It("should return predicated value", func() {
			act()

			Expect(observedValue).To(Equal(predicatedValue))
			Expect(observedError).To(BeNil())
		})
	})
})

var _ = DescribeFunction(fn.SafeCallable, func() {
	var input typing.Callable
	var observedError error

	act := func() {
		observedError = fn.SafeCallable(input)
	}

	BeforeEach(func() {
		observedError = nil
	})

	Context("with panicking callable", func() {
		panickedError := errors.New("something went wrong")

		BeforeEach(func() {
			input = func() {
				panic(panickedError)
			}
		})

		It("should return error", func() {
			act()

			Expect(observedError).To(Equal(panickedError))
		})
	})

	Context("with non panicking callable", func() {
		BeforeEach(func() {
			input = func() {}
		})

		It("should return no error", func() {
			act()

			Expect(observedError).To(BeNil())
		})
	})
})
