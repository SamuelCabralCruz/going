//go:build test

package result_test

import (
	"errors"
	"fmt"
	"github.com/SamuelCabralCruz/going/fn/result"
	"github.com/SamuelCabralCruz/going/fn/tuple/validation"
	"github.com/SamuelCabralCruz/going/fn/typing"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = DescribeType[result.Result[any]](func() {
	DescribeFunction(result.Ok[any], func() {
		var input any
		var observed result.Result[any]

		act := func() {
			observed = result.Ok[any](input)
		}

		BeforeEach(func() {
			input = "some value"
		})

		It("should return ok result", func() {
			act()

			Expect(observed.IsOk()).To(BeTrue())
			Expect(observed.OrEmpty()).To(Equal(input))
			Expect(observed.IsError()).To(BeFalse())
			Expect(observed.Error()).To(BeNil())
		})
	})

	DescribeFunction(result.Error[any], func() {
		var input error
		var observed result.Result[any]

		act := func() {
			observed = result.Error[any](input)
		}

		BeforeEach(func() {
			input = errors.New("something went wrong")
		})

		It("should return error result", func() {
			act()

			Expect(observed.IsOk()).To(BeFalse())
			Expect(observed.OrEmpty()).To(BeZero())
			Expect(observed.IsError()).To(BeTrue())
			Expect(observed.Error()).To(Equal(input))
		})
	})

	DescribeFunction(result.Errorf[any], func() {
		var input string
		var observed result.Result[any]

		act := func() {
			observed = result.Errorf[any](input)
		}

		BeforeEach(func() {
			input = "something went wrong"
		})

		It("should return error result", func() {
			act()

			Expect(observed.IsOk()).To(BeFalse())
			Expect(observed.OrEmpty()).To(BeZero())
			Expect(observed.IsError()).To(BeTrue())
			Expect(observed.Error().Error()).To(Equal(input))
		})
	})

	DescribeFunction(result.FromAssertion[any], func() {
		var inputValue any
		var inputError error
		var observed result.Result[any]

		act := func() {
			observed = result.FromAssertion(inputValue, inputError)
		}

		BeforeEach(func() {
			inputValue = nil
			inputError = nil
		})

		Context("with error assertion", func() {
			BeforeEach(func() {
				inputError = errors.New("something went wrong")
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(inputError))
			})
		})

		Context("without error assertion", func() {
			BeforeEach(func() {
				inputValue = "some value"
			})

			It("should return ok result", func() {
				act()

				Expect(observed.IsOk()).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
				Expect(observed.IsError()).To(BeFalse())
				Expect(observed.Error()).To(BeNil())
			})
		})
	})

	DescribeFunction(result.FromValidation[any], func() {
		var inputValue any
		var inputOk bool
		var observed result.Result[any]

		act := func() {
			observed = result.FromValidation(inputValue, inputOk)
		}

		BeforeEach(func() {
			inputValue = nil
			inputOk = false
		})

		Context("with not ok validation", func() {
			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(BeAssignableToTypeOf(validation.InvalidValueError{}))
			})
		})

		Context("with ok validation", func() {
			BeforeEach(func() {
				inputOk = true
				inputValue = "some value"
			})

			It("should return ok result", func() {
				act()

				Expect(observed.IsOk()).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
				Expect(observed.IsError()).To(BeFalse())
				Expect(observed.Error()).To(BeNil())
			})
		})
	})

	DescribeFunction(result.FromSupplier[any], func() {
		var input typing.Supplier[any]
		var observed result.Result[any]

		act := func() {
			observed = result.FromSupplier(input)
		}

		Context("with panicking supplier", func() {
			panickedError := errors.New("something went wrong")

			BeforeEach(func() {
				input = func() any {
					panic(panickedError)
				}
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(panickedError))
			})
		})

		Context("with non panicking supplier", func() {
			inputValue := "some value"

			BeforeEach(func() {
				input = func() any {
					return inputValue
				}
			})

			It("should return ok result", func() {
				act()

				Expect(observed.IsOk()).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
				Expect(observed.IsError()).To(BeFalse())
				Expect(observed.Error()).To(BeNil())
			})
		})
	})

	DescribeFunction(result.FromProducer[any], func() {
		var input typing.Producer[any]
		var observed result.Result[any]

		act := func() {
			observed = result.FromProducer(input)
		}

		Context("with panicking producer", func() {
			panickedError := errors.New("something went wrong")

			BeforeEach(func() {
				input = func() (any, error) {
					panic(panickedError)
				}
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(panickedError))
			})
		})

		Context("with producer returning an error", func() {
			panickedError := errors.New("something went wrong")

			BeforeEach(func() {
				input = func() (any, error) {
					return nil, panickedError
				}
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(panickedError))
			})
		})

		Context("with non panicking producer", func() {
			inputValue := "some value"

			BeforeEach(func() {
				input = func() (any, error) {
					return inputValue, nil
				}
			})

			It("should return ok result", func() {
				act()

				Expect(observed.IsOk()).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
				Expect(observed.IsError()).To(BeFalse())
				Expect(observed.Error()).To(BeNil())
			})
		})
	})

	var value any
	var cut result.Result[any]

	BeforeEach(func() {
		value = "some value"
		cut = result.Ok[any](value)
	})

	DescribeFunction(cut.Get, func() {
		var observedValue any
		var observedError error

		act := func() {
			observedValue, observedError = cut.Get()
		}

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should return an error", func() {
				act()

				Expect(observedValue).To(BeNil())
				Expect(observedError).NotTo(BeNil())
				Expect(observedError).To(Equal(err))
			})
		})

		Context("with ok result", func() {
			It("should return the value", func() {
				act()

				Expect(observedValue).NotTo(BeNil())
				Expect(observedError).To(BeNil())
				Expect(observedValue).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.GetOrPanic, func() {
		var observed any

		act := func() {
			observed = cut.GetOrPanic()
		}

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should panic", func() {
				Expect(act).To(PanicWith(err))
			})
		})

		Context("with ok result", func() {
			It("should return the value", func() {
				act()

				Expect(observed).NotTo(BeNil())
				Expect(observed).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.GetOrPanicWith, func() {
		var input error
		var observed any

		act := func() {
			observed = cut.GetOrPanicWith(input)
		}

		BeforeEach(func() {
			input = errors.New("something went wrong")
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should panic", func() {
				Expect(act).To(PanicWith(input))
			})
		})

		Context("with ok result", func() {
			It("should return the value", func() {
				act()

				Expect(observed).NotTo(BeNil())
				Expect(observed).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.OrElseGet, func() {
		var input typing.Supplier[any]
		var observed any
		suppliedValue := "some supplied value"

		act := func() {
			observed = cut.OrElseGet(input)
		}

		BeforeEach(func() {
			input = func() any {
				return suppliedValue
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should return supplied value", func() {
				act()

				Expect(observed).To(Equal(suppliedValue))
			})

			Context("with panicking supplier", func() {
				BeforeEach(func() {
					input = func() any {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return zero value", func() {
					act()

					Expect(observed).To(BeZero())
				})
			})
		})

		Context("with ok result", func() {
			It("should return value", func() {
				act()

				Expect(observed).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.OrElse, func() {
		var input any
		var observed any

		act := func() {
			observed = cut.OrElse(input)
		}

		BeforeEach(func() {
			input = "some fallback value"
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should return fallback value", func() {
				act()

				Expect(observed).To(Equal(input))
			})
		})

		Context("with ok result", func() {
			It("should return value", func() {
				act()

				Expect(observed).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.FlatMap, func() {
		var cut result.Result[string]
		var input typing.Mapper[result.Result[string]]
		var observed result.Result[string]

		act := func() {
			observed = cut.FlatMap(input)
		}

		BeforeEach(func() {
			input = func(o result.Result[string]) result.Result[string] {
				return result.Ok(strings.TrimPrefix(o.OrEmpty(), "prefix"))
			}
			cut = result.Ok("prefixSomeValue")
		})

		It("should map", func() {
			act()

			Expect(observed.OrEmpty()).To(Equal("SomeValue"))
		})

		Context("with panicking mapper", func() {
			panickedError := errors.New("something went wrong")

			BeforeEach(func() {
				input = func(_ result.Result[string]) result.Result[string] {
					panic(panickedError)
				}
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(panickedError))
			})
		})
	})

	DescribeFunction(cut.Map, func() {
		var cut result.Result[string]
		var input typing.Mapper[string]
		var observed result.Result[string]
		var invoked bool

		act := func() {
			observed = cut.Map(input)
		}

		BeforeEach(func() {
			invoked = false
			input = func(s string) string {
				invoked = true
				return strings.TrimPrefix(s, "prefix")
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[string](err)
			})

			It("should not map", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})

		Context("with ok result", func() {
			BeforeEach(func() {
				cut = result.Ok("prefixSomeValue")
			})

			It("should map", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal("SomeValue"))
			})

			Context("with panicking mapper", func() {
				panickedError := errors.New("something went wrong")

				BeforeEach(func() {
					input = func(_ string) string {
						panic(panickedError)
					}
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(Equal(panickedError))
				})
			})
		})
	})

	DescribeFunction(cut.MapError, func() {
		var cut result.Result[string]
		var input typing.Transformer[error, string]
		var observed result.Result[string]
		var invoked bool
		var received error
		suppliedValue := "some fallback value"

		act := func() {
			observed = cut.MapError(input)
		}

		BeforeEach(func() {
			invoked = false
			received = nil
			input = func(err error) string {
				invoked = true
				received = err
				return suppliedValue
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[string](err)
			})

			It("should transform error into value", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(received).To(Equal(err))
				Expect(observed.OrEmpty()).To(Equal(suppliedValue))
			})

			Context("with panicking transformer", func() {
				panickedError := errors.New("something went wrong")

				BeforeEach(func() {
					input = func(_ error) string {
						panic(panickedError)
					}
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(Equal(panickedError))
				})
			})
		})

		Context("with ok result", func() {
			BeforeEach(func() {
				cut = result.Ok("some value")
			})

			It("should not transform", func() {
				act()

				Expect(invoked).To(BeFalse())
				Expect(observed.OrEmpty()).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.SwitchMap, func() {
		var cut result.Result[string]
		var inputOnPresent typing.Mapper[string]
		var inputOnError typing.Transformer[error, string]
		var observed result.Result[string]

		act := func() {
			observed = cut.SwitchMap(inputOnPresent, inputOnError)
		}

		BeforeEach(func() {
			inputOnPresent = func(s string) string {
				return fmt.Sprintf("on present - %s", s)
			}
			inputOnError = func(err error) string {
				return fmt.Sprintf("on error - %s", err.Error())
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[string](err)
			})

			It("should transform error into value", func() {
				act()

				Expect(observed.OrEmpty()).To(Equal("on error - something went wrong"))
			})

			Context("with panicking transformer", func() {
				panickedError := errors.New("something went wrong")

				BeforeEach(func() {
					inputOnError = func(_ error) string {
						panic(panickedError)
					}
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(Equal(panickedError))
				})
			})
		})

		Context("with ok result", func() {
			BeforeEach(func() {
				cut = result.Ok("some value")
			})

			It("should map value", func() {
				act()

				Expect(observed.OrEmpty()).To(Equal("on present - some value"))
			})

			Context("with panicking mapper", func() {
				panickedError := errors.New("something went wrong")

				BeforeEach(func() {
					inputOnPresent = func(_ string) string {
						panic(panickedError)
					}
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(Equal(panickedError))
				})
			})
		})
	})

	DescribeFunction(cut.IfOk, func() {
		var input typing.Consumer[any]
		var received any
		var invoked bool

		act := func() {
			cut.IfOk(input)
		}

		BeforeEach(func() {
			invoked = false
			received = nil
			input = func(r any) {
				invoked = true
				received = r
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should not invoke consumer", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})

		Context("with ok result", func() {
			It("should invoke consumer", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(received).To(Equal(value))
			})

			Context("with panicking consumer", func() {
				BeforeEach(func() {
					input = func(_ any) {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})
	})

	DescribeFunction(cut.IfError, func() {
		var input typing.Consumer[error]
		var received error
		var invoked bool

		act := func() {
			cut.IfError(input)
		}

		BeforeEach(func() {
			invoked = false
			received = nil
			input = func(err error) {
				invoked = true
				received = err
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should invoke consumer", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(received).To(Equal(err))
			})

			Context("with panicking consumer", func() {
				BeforeEach(func() {
					input = func(_ error) {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})

		Context("with ok result", func() {
			It("should invoke consumer", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})
	})

	DescribeFunction(cut.Switch, func() {
		var inputOnOk typing.Consumer[any]
		var inputOnError typing.Consumer[error]
		var invokedOnOk bool
		var receivedOnOk any
		var invokedOnError bool
		var receivedOnError error

		act := func() {
			cut.Switch(inputOnOk, inputOnError)
		}

		BeforeEach(func() {
			inputOnOk = func(r any) {
				invokedOnOk = true
				receivedOnOk = r
			}
			inputOnError = func(err error) {
				invokedOnError = true
				receivedOnError = err
			}
			invokedOnOk = false
			receivedOnOk = nil
			invokedOnError = false
			receivedOnError = nil
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[any](err)
			})

			It("should invoke on error", func() {
				act()

				Expect(invokedOnOk).To(BeFalse())
				Expect(receivedOnOk).To(BeNil())
				Expect(invokedOnError).To(BeTrue())
				Expect(receivedOnError).To(Equal(err))
			})

			Context("with panicking on error", func() {
				BeforeEach(func() {
					inputOnError = func(_ error) {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})

		Context("with ok result", func() {
			It("should invoke on ok", func() {
				act()

				Expect(invokedOnOk).To(BeTrue())
				Expect(receivedOnOk).To(Equal(value))
				Expect(invokedOnError).To(BeFalse())
				Expect(receivedOnError).To(BeNil())
			})

			Context("with panicking on ok", func() {
				BeforeEach(func() {
					inputOnOk = func(_ any) {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})
	})

	DescribeFunction(cut.Filter, func() {
		var cut result.Result[string]
		var input typing.Predicate[string]
		var observed result.Result[string]

		act := func() {
			observed = cut.Filter(input)
		}

		BeforeEach(func() {
			input = func(s string) bool {
				return len(s) == 3
			}
		})

		Context("with error result", func() {
			err := errors.New("something went wrong")

			BeforeEach(func() {
				cut = result.Error[string](err)
			})

			It("should return error result", func() {
				act()

				Expect(observed.IsOk()).To(BeFalse())
				Expect(observed.OrEmpty()).To(BeZero())
				Expect(observed.IsError()).To(BeTrue())
				Expect(observed.Error()).To(Equal(err))
			})
		})

		Context("with ok result", func() {
			Context("with non matching value", func() {
				BeforeEach(func() {
					cut = result.Ok("some")
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(BeAssignableToTypeOf(result.FilteredValueError{}))
				})
			})

			Context("with matching value", func() {
				BeforeEach(func() {
					cut = result.Ok("som")
				})

				It("should return result", func() {
					act()

					Expect(observed.IsOk()).To(BeTrue())
					Expect(observed.IsError()).To(BeFalse())
					Expect(observed).To(Equal(cut))
				})
			})

			Context("with panicking predicate", func() {
				panickedError := errors.New("something went wrong")

				BeforeEach(func() {
					input = func(_ string) bool {
						panic(panickedError)
					}
				})

				It("should return error result", func() {
					act()

					Expect(observed.IsOk()).To(BeFalse())
					Expect(observed.OrEmpty()).To(BeZero())
					Expect(observed.IsError()).To(BeTrue())
					Expect(observed.Error()).To(Equal(panickedError))
				})
			})
		})
	})
})
