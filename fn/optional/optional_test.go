//go:build test

package optional_test

import (
	"errors"
	"fmt"
	"github.com/SamuelCabralCruz/going/fn/optional"
	"github.com/SamuelCabralCruz/going/fn/typing"
	. "github.com/SamuelCabralCruz/going/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = DescribeType[optional.Optional[any]](func() {
	DescribeFunction(optional.Empty[any], func() {
		var observed optional.Optional[any]

		act := func() {
			observed = optional.Empty[any]()
		}

		It("should return empty optional", func() {
			act()

			Expect(observed.IsPresent()).To(BeFalse())
			Expect(observed.IsAbsent()).To(BeTrue())
			Expect(observed.OrEmpty()).To(BeZero())
		})
	})

	DescribeFunction(optional.Of[any], func() {
		var input any
		var observed optional.Optional[any]

		act := func() {
			observed = optional.Of(input)
		}

		DescribeTable("should return present optional", func(tableInput any) {
			input = tableInput

			act()

			Expect(observed.IsPresent()).To(BeTrue())
			Expect(observed.IsAbsent()).To(BeFalse())
			if input == nil {
				Expect(observed.OrEmpty()).To(BeNil())
			} else {
				Expect(observed.OrEmpty()).To(Equal(input))
			}
		},
			CreateTableEntries(
				[]string{"input"},
				[]any{"some value"},
				[]any{""},
				[]any{1},
				[]any{12345},
				[]any{0},
				[]any{nil},
			),
		)
	})

	DescribeFunction(optional.OfNullable[any], func() {
		var input any
		var observed optional.Optional[any]

		act := func() {
			observed = optional.OfNullable(input)
		}

		Context("with zero value", func() {
			DescribeTable("should return empty optional", func(tableInput any) {
				input = tableInput

				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			},
				CreateTableEntries(
					[]string{"input"},
					[]any{nil},
					[]any{0},
					[]any{""},
					[]any{[]any{}},
					[]any{map[any]any{}},
				),
			)
		})

		Context("with non zero value", func() {
			DescribeTable("should return present optional", func(tableInput any) {
				input = tableInput

				act()

				Expect(observed.IsPresent()).To(BeTrue())
				Expect(observed.IsAbsent()).To(BeFalse())
				Expect(observed.OrEmpty()).To(Equal(input))
			},
				CreateTableEntries(
					[]string{"input"},
					[]any{"some value"},
					[]any{1},
				),
			)
		})
	})

	DescribeFunction(optional.FromAssertion[any], func() {
		var inputValue any
		var inputError error
		var observed optional.Optional[any]

		act := func() {
			observed = optional.FromAssertion(inputValue, inputError)
		}

		BeforeEach(func() {
			inputValue = nil
			inputError = nil
		})

		Context("with error assertion", func() {
			BeforeEach(func() {
				inputError = errors.New("something went wrong")
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("without error assertion", func() {
			Context("with zero value", func() {
				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})

			Context("with non zero value", func() {
				BeforeEach(func() {
					inputValue = "some value"
				})

				It("should return present optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeTrue())
					Expect(observed.IsAbsent()).To(BeFalse())
					Expect(observed.OrEmpty()).To(Equal(inputValue))
				})
			})
		})
	})

	DescribeFunction(optional.FromValidation[any], func() {
		var inputValue any
		var inputOk bool
		var observed optional.Optional[any]

		act := func() {
			observed = optional.FromValidation(inputValue, inputOk)
		}

		BeforeEach(func() {
			inputValue = nil
			inputOk = false
		})

		Context("with not ok validation", func() {
			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with ok validation", func() {
			BeforeEach(func() {
				inputOk = true
			})

			Context("with zero value", func() {
				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})

			Context("with non zero value", func() {
				BeforeEach(func() {
					inputValue = "some value"
				})

				It("should return present optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeTrue())
					Expect(observed.IsAbsent()).To(BeFalse())
					Expect(observed.OrEmpty()).To(Equal(inputValue))
				})
			})
		})
	})

	DescribeFunction(optional.FromSupplier[any], func() {
		var input typing.Supplier[any]
		var observed optional.Optional[any]

		act := func() {
			observed = optional.FromSupplier(input)
		}

		Context("with panicking supplier", func() {
			BeforeEach(func() {
				input = func() any {
					panic(errors.New("something went wrong"))
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with supplier returning zero value", func() {
			BeforeEach(func() {
				input = func() any {
					return nil
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with supplier returning non zero value", func() {
			inputValue := "some value"

			BeforeEach(func() {
				input = func() any {
					return inputValue
				}
			})

			It("should return present optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeTrue())
				Expect(observed.IsAbsent()).To(BeFalse())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
			})
		})
	})

	DescribeFunction(optional.FromProducer[any], func() {
		var input typing.Producer[any]
		var observed optional.Optional[any]

		act := func() {
			observed = optional.FromProducer(input)
		}

		Context("with panicking producer", func() {
			BeforeEach(func() {
				input = func() (any, error) {
					panic(errors.New("something went wrong"))
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with producer returning an error", func() {
			BeforeEach(func() {
				input = func() (any, error) {
					return nil, errors.New("something went wrong")
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with producer returning zero value", func() {
			BeforeEach(func() {
				input = func() (any, error) {
					return nil, nil
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with producer returning non zero value", func() {
			inputValue := "some value"

			BeforeEach(func() {
				input = func() (any, error) {
					return inputValue, nil
				}
			})

			It("should return present optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeTrue())
				Expect(observed.IsAbsent()).To(BeFalse())
				Expect(observed.OrEmpty()).To(Equal(inputValue))
			})
		})
	})

	var cut optional.Optional[any]

	BeforeEach(func() {
		cut = optional.Empty[any]()
	})

	DescribeFunction(cut.Get, func() {
		var observedValue any
		var observedError error

		act := func() {
			observedValue, observedError = cut.Get()
		}

		Context("with empty optional", func() {
			It("should return an error", func() {
				act()

				Expect(observedValue).To(BeNil())
				Expect(observedError).NotTo(BeNil())
				Expect(observedError).To(BeAssignableToTypeOf(optional.NoSuchElementError{}))
			})
		})

		Context("with present optional", func() {
			inputValue := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](inputValue)
			})

			It("should return the value", func() {
				act()

				Expect(observedValue).NotTo(BeNil())
				Expect(observedError).To(BeNil())
				Expect(observedValue).To(Equal(inputValue))
			})
		})
	})

	DescribeFunction(cut.GetOrPanic, func() {
		var observed any

		act := func() {
			observed = cut.GetOrPanic()
		}

		Context("with empty optional", func() {
			It("should panic", func() {
				Expect(act).To(PanicWith(BeAssignableToTypeOf(optional.NoSuchElementError{})))
			})
		})

		Context("with present optional", func() {
			inputValue := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](inputValue)
			})

			It("should return the value", func() {
				act()

				Expect(observed).NotTo(BeNil())
				Expect(observed).To(Equal(inputValue))
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

		Context("with empty optional", func() {
			It("should panic", func() {
				Expect(act).To(PanicWith(input))
			})
		})

		Context("with present optional", func() {
			inputValue := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](inputValue)
			})

			It("should return the value", func() {
				act()

				Expect(observed).NotTo(BeNil())
				Expect(observed).To(Equal(inputValue))
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

		Context("with empty optional", func() {
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

		Context("with present optional", func() {
			value := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](value)
			})

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

		Context("with empty optional", func() {
			It("should return fallback value", func() {
				act()

				Expect(observed).To(Equal(input))
			})
		})

		Context("with present optional", func() {
			value := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](value)
			})

			It("should return value", func() {
				act()

				Expect(observed).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.FlatMap, func() {
		var cut optional.Optional[string]
		var input typing.Mapper[optional.Optional[string]]
		var observed optional.Optional[string]

		act := func() {
			observed = cut.FlatMap(input)
		}

		BeforeEach(func() {
			input = func(o optional.Optional[string]) optional.Optional[string] {
				return optional.Of(strings.TrimPrefix(o.OrEmpty(), "prefix"))
			}
			cut = optional.Of("prefixSomeValue")
		})

		It("should map", func() {
			act()

			Expect(observed.OrEmpty()).To(Equal("SomeValue"))
		})

		Context("with panicking mapper", func() {
			BeforeEach(func() {
				input = func(_ optional.Optional[string]) optional.Optional[string] {
					panic(errors.New("something went wrong"))
				}
			})

			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})
	})

	DescribeFunction(cut.Map, func() {
		var cut optional.Optional[string]
		var input typing.Mapper[string]
		var observed optional.Optional[string]
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

		Context("with empty optional", func() {
			It("should not map", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})

		Context("with present optional", func() {
			BeforeEach(func() {
				cut = optional.Of("prefixSomeValue")
			})

			It("should map", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal("SomeValue"))
			})

			Context("with panicking mapper", func() {
				BeforeEach(func() {
					input = func(_ string) string {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})
		})
	})

	DescribeFunction(cut.MapEmpty, func() {
		var cut optional.Optional[string]
		var input typing.Supplier[string]
		var observed optional.Optional[string]
		var invoked bool
		suppliedValue := "some fallback value"

		act := func() {
			observed = cut.MapEmpty(input)
		}

		BeforeEach(func() {
			invoked = false
			input = func() string {
				invoked = true
				return suppliedValue
			}
		})

		Context("with empty optional", func() {
			It("should supply value", func() {
				act()

				Expect(invoked).To(BeTrue())
				Expect(observed.OrEmpty()).To(Equal(suppliedValue))
			})

			Context("with panicking supplier", func() {
				BeforeEach(func() {
					input = func() string {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})
		})

		Context("with present optional", func() {
			value := "some value"

			BeforeEach(func() {
				cut = optional.Of(value)
			})

			It("should not supply", func() {
				act()

				Expect(invoked).To(BeFalse())
				Expect(observed.OrEmpty()).To(Equal(value))
			})
		})
	})

	DescribeFunction(cut.SwitchMap, func() {
		var cut optional.Optional[string]
		var inputOnPresent typing.Mapper[string]
		var inputOnEmpty typing.Supplier[string]
		var observed optional.Optional[string]

		act := func() {
			observed = cut.SwitchMap(inputOnPresent, inputOnEmpty)
		}

		BeforeEach(func() {
			inputOnPresent = func(s string) string {
				return fmt.Sprintf("on present - %s", s)
			}
			inputOnEmpty = func() string {
				return "on empty"
			}
		})

		Context("with empty optional", func() {
			It("should supply value", func() {
				act()

				Expect(observed.OrEmpty()).To(Equal("on empty"))
			})

			Context("with panicking supplier", func() {
				BeforeEach(func() {
					inputOnEmpty = func() string {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})
		})

		Context("with present optional", func() {
			BeforeEach(func() {
				cut = optional.Of("some value")
			})

			It("should map value", func() {
				act()

				Expect(observed.OrEmpty()).To(Equal("on present - some value"))
			})

			Context("with panicking mapper", func() {
				BeforeEach(func() {
					inputOnPresent = func(_ string) string {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})
		})
	})

	DescribeFunction(cut.IfPresent, func() {
		var input typing.Consumer[any]
		var received any
		var invoked bool

		act := func() {
			cut.IfPresent(input)
		}

		BeforeEach(func() {
			invoked = false
			received = nil
			input = func(r any) {
				invoked = true
				received = r
			}
		})

		Context("with empty optional", func() {
			It("should not invoke consumer", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})

		Context("with present optional", func() {
			value := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](value)
			})

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

	DescribeFunction(cut.IfAbsent, func() {
		var input typing.Callable
		var invoked bool

		act := func() {
			cut.IfAbsent(input)
		}

		BeforeEach(func() {
			invoked = false
			input = func() {
				invoked = true
			}
		})

		Context("with empty optional", func() {
			It("should invoke callable", func() {
				act()

				Expect(invoked).To(BeTrue())
			})

			Context("with panicking callable", func() {
				BeforeEach(func() {
					input = func() {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})

		Context("with present optional", func() {
			BeforeEach(func() {
				cut = optional.Of[any]("some value")
			})

			It("should not invoke callable", func() {
				act()

				Expect(invoked).To(BeFalse())
			})
		})
	})

	DescribeFunction(cut.Switch, func() {
		var inputOnPresent typing.Consumer[any]
		var inputOnAbsent typing.Callable
		var invokedOnPresent bool
		var receivedOnPresent any
		var invokedOnAbsent bool

		act := func() {
			cut.Switch(inputOnPresent, inputOnAbsent)
		}

		BeforeEach(func() {
			inputOnPresent = func(r any) {
				invokedOnPresent = true
				receivedOnPresent = r
			}
			inputOnAbsent = func() {
				invokedOnAbsent = true
			}
			invokedOnPresent = false
			receivedOnPresent = nil
			invokedOnAbsent = false
		})

		Context("with empty optional", func() {
			It("should invoke on absent", func() {
				act()

				Expect(invokedOnPresent).To(BeFalse())
				Expect(invokedOnAbsent).To(BeTrue())
			})

			Context("with panicking on absent", func() {
				BeforeEach(func() {
					inputOnAbsent = func() {
						panic(errors.New("something went wrong"))
					}
				})

				It("should panic", func() {
					Expect(act).To(Panic())
				})
			})
		})

		Context("with present optional", func() {
			value := "some value"

			BeforeEach(func() {
				cut = optional.Of[any](value)
			})

			It("should invoke on present", func() {
				act()

				Expect(invokedOnPresent).To(BeTrue())
				Expect(invokedOnAbsent).To(BeFalse())
				Expect(receivedOnPresent).To(Equal(value))
			})

			Context("with panicking on present", func() {
				BeforeEach(func() {
					inputOnPresent = func(_ any) {
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
		var cut optional.Optional[string]
		var input typing.Predicate[string]
		var observed optional.Optional[string]

		act := func() {
			observed = cut.Filter(input)
		}

		BeforeEach(func() {
			input = func(s string) bool {
				return len(s) == 3
			}
		})

		Context("with empty optional", func() {
			It("should return empty optional", func() {
				act()

				Expect(observed.IsPresent()).To(BeFalse())
				Expect(observed.IsAbsent()).To(BeTrue())
				Expect(observed.OrEmpty()).To(BeZero())
			})
		})

		Context("with present optional", func() {
			Context("with non matching value", func() {
				BeforeEach(func() {
					cut = optional.Of("some")
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})

			Context("with matching value", func() {
				BeforeEach(func() {
					cut = optional.Of("som")
				})

				It("should return optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeTrue())
					Expect(observed.IsAbsent()).To(BeFalse())
					Expect(observed).To(Equal(cut))
				})
			})

			Context("with panicking predicate", func() {
				BeforeEach(func() {
					input = func(_ string) bool {
						panic(errors.New("something went wrong"))
					}
				})

				It("should return empty optional", func() {
					act()

					Expect(observed.IsPresent()).To(BeFalse())
					Expect(observed.IsAbsent()).To(BeTrue())
					Expect(observed.OrEmpty()).To(BeZero())
				})
			})
		})
	})
})
