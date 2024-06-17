//go:build test

package assertion_test

import (
	"errors"
	"github.com/SamuelCabralCruz/going/fn/tuple/assertion"
	"github.com/SamuelCabralCruz/going/fn/typing"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(assertion.IgnoreError[any], func() {
	var inputValue any
	var inputError error
	var observed any

	act := func() {
		observed = assertion.IgnoreError(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputError = errors.New("something went wrong")
	})

	It("should return asserted value", func() {
		act()

		Expect(observed).To(Equal(inputValue))
	})
})

var _ = DescribeFunction(assertion.IgnoreValue[any], func() {
	var inputValue any
	var inputError error
	var observed any

	act := func() {
		observed = assertion.IgnoreValue(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputError = errors.New("something went wrong")
	})

	It("should return asserted error", func() {
		act()

		Expect(observed).To(Equal(inputError))
	})
})

var _ = DescribeFunction(assertion.GetOrEmpty[any], func() {
	var inputValue any
	var inputError error
	var observed any

	act := func() {
		observed = assertion.GetOrEmpty(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = nil
		inputError = nil
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			inputError = errors.New("something went wrong")
		})

		It("should return empty value", func() {
			act()

			Expect(observed).To(BeZero())
		})
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			inputValue = "some value"
		})

		It("should return asserted value", func() {
			act()

			Expect(observed).To(Equal(inputValue))
		})
	})
})

var _ = DescribeFunction(assertion.GetOrPanic[any], func() {
	var inputValue any
	var inputError error
	var observed any

	act := func() {
		observed = assertion.GetOrPanic(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = nil
		inputError = nil
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			inputError = errors.New("something went wrong")
		})

		It("should panic", func() {
			Expect(act).To(PanicWith(inputError))
		})
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			inputValue = "some value"
		})

		It("should return asserted value", func() {
			act()

			Expect(observed).To(Equal(inputValue))
		})
	})
})

var _ = DescribeFunction(assertion.PanicIfError[any], func() {
	var inputValue any
	var inputError error

	act := func() {
		assertion.PanicIfError(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = nil
		inputError = nil
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			inputError = errors.New("something went wrong")
		})

		It("should panic", func() {
			Expect(act).To(PanicWith(inputError))
		})
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			inputValue = "some value"
		})

		It("should not panic", func() {
			Expect(act).NotTo(Panic())
		})
	})
})

var _ = DescribeFunction(assertion.FromValue[any], func() {
	var input any
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = assertion.FromValue(input)
	}

	BeforeEach(func() {
		input = "some value"
	})

	It("should return ok assertion", func() {
		act()

		Expect(observedValue).To(Equal(input))
		Expect(observedError).To(BeNil())
	})
})

var _ = DescribeFunction(assertion.FromError[any], func() {
	var input error
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = assertion.FromError[any](input)
	}

	BeforeEach(func() {
		input = errors.New("something went wrong")
	})

	It("should return error assertion", func() {
		act()

		Expect(observedValue).To(BeZero())
		Expect(observedError).To(Equal(input))
	})
})

var _ = DescribeFunction(assertion.FromReversed[any], func() {
	var inputValue any
	var inputError error
	var observedValue any
	var observedError error

	act := func() {
		observedValue, observedError = assertion.FromReversed[any](inputError, inputValue)
	}

	BeforeEach(func() {
		inputValue = "some value"
		inputError = errors.New("something went wrong")
	})

	It("should return corresponding assertion", func() {
		act()

		Expect(observedValue).To(Equal(inputValue))
		Expect(observedError).To(Equal(inputError))
	})
})

var _ = DescribeFunction(assertion.ToValidation[any], func() {
	var inputValue any
	var inputError error
	var observedValue any
	var observedOk bool

	act := func() {
		observedValue, observedOk = assertion.ToValidation(inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = nil
		inputError = nil
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			inputValue = "some value"
		})

		It("should return ok validation", func() {
			act()

			Expect(observedValue).To(Equal(inputValue))
			Expect(observedOk).To(BeTrue())
		})
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			inputValue = "some ignored value"
			inputError = errors.New("something went wrong")
		})

		It("should return non ok validation", func() {
			act()

			Expect(observedValue).To(BeZero())
			Expect(observedOk).To(BeFalse())
		})
	})
})

var _ = DescribeFunction(assertion.Switch[any, any], func() {
	var inputValue string
	var inputError error
	var receivedOnOk string
	var receivedOnError error
	var transformedOnOk int
	var transformedOnError int
	var onOk typing.Transformer[string, int]
	var onError typing.Transformer[error, int]
	var observed func(typing.Transformer[string, int], typing.Transformer[error, int]) int

	act := func() {
		observed = assertion.Switch[string, int](inputValue, inputError)
	}

	BeforeEach(func() {
		inputValue = ""
		inputError = nil
		receivedOnOk = ""
		receivedOnError = nil
		transformedOnOk = 1234
		transformedOnError = 4321
		onOk = func(v string) int {
			receivedOnOk = v
			return transformedOnOk
		}
		onError = func(err error) int {
			receivedOnError = err
			return transformedOnError
		}
	})

	Context("with ok assertion", func() {
		BeforeEach(func() {
			inputValue = "some value"
		})

		It("should transform asserted value using on ok transformer", func() {
			act()

			Expect(observed(onOk, onError)).To(Equal(transformedOnOk))
			Expect(receivedOnOk).To(Equal(inputValue))
		})
	})

	Context("with error assertion", func() {
		BeforeEach(func() {
			inputError = errors.New("something went wrong")
		})

		It("should transform asserted error using on error transformer", func() {
			act()

			Expect(observed(onOk, onError)).To(Equal(transformedOnError))
			Expect(receivedOnError).To(Equal(inputError))
		})
	})
})
