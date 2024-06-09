//go:build test

package result_test

import (
	"errors"
	"github.com/SamuelCabralCruz/went/fn/result"
	"github.com/SamuelCabralCruz/went/fn/typing"
	. "github.com/SamuelCabralCruz/went/kinggo"
	"github.com/SamuelCabralCruz/went/roar"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(result.Transform[any, any], func() {
	var cut result.Result[string]
	var transformer typing.Transformer[string, int]
	var observed result.Result[int]

	act := func() {
		observed = result.Transform(cut, transformer)
	}

	BeforeEach(func() {
		transformer = func(s string) int {
			return len(s)
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
			Expect(observed.IsError()).To(BeTrue())
			Expect(observed.OrEmpty()).To(BeZero())
			Expect(observed.Error()).To(Equal(err))
		})
	})

	Context("with ok result", func() {
		BeforeEach(func() {
			cut = result.Ok("some value")
		})

		It("should return transformed optional", func() {
			act()

			Expect(observed.IsOk()).To(BeTrue())
			Expect(observed.IsError()).To(BeFalse())
			Expect(observed.OrEmpty()).To(Equal(10))
			Expect(observed.Error()).To(BeNil())
		})
	})
})

var _ = DescribeFunction(result.FilterOk[any], func() {
	var inputs []result.Result[any]
	var observed []any

	act := func() {
		observed = result.FilterOk(inputs...)
	}

	BeforeEach(func() {
		inputs = []result.Result[any]{
			result.Error[any](errors.New("something went wrong 1")),
			result.Error[any](errors.New("something went wrong 2")),
			result.Error[any](errors.New("something went wrong 3")),
			result.Ok[any]("a"),
			result.Error[any](errors.New("something went wrong 4")),
			result.Ok[any]("b"),
			result.Ok[any]("c"),
			result.Error[any](errors.New("something went wrong 5")),
			result.Error[any](errors.New("something went wrong 6")),
			result.Error[any](errors.New("something went wrong 7")),
			result.Error[any](errors.New("something went wrong 8")),
			result.Ok[any]("d"),
			result.Error[any](errors.New("something went wrong 9")),
			result.Ok[any]("e"),
			result.Error[any](errors.New("something went wrong 10")),
			result.Error[any](errors.New("something went wrong 11")),
			result.Ok[any]("f"),
			result.Error[any](errors.New("something went wrong 12")),
			result.Error[any](errors.New("something went wrong 13")),
			result.Error[any](errors.New("something went wrong 14")),
		}
	})

	It("should return ok values", func() {
		act()

		Expect(observed).To(Equal([]any{"a", "b", "c", "d", "e", "f"}))
	})
})

var _ = DescribeFunction(result.FilterError[any], func() {
	var inputs []result.Result[any]
	var observed []error

	act := func() {
		observed = result.FilterError(inputs...)
	}

	BeforeEach(func() {
		inputs = []result.Result[any]{
			result.Error[any](errors.New("something went wrong 1")),
			result.Error[any](errors.New("something went wrong 2")),
			result.Error[any](errors.New("something went wrong 3")),
			result.Ok[any]("a"),
			result.Error[any](errors.New("something went wrong 4")),
			result.Ok[any]("b"),
			result.Ok[any]("c"),
			result.Error[any](errors.New("something went wrong 5")),
			result.Error[any](errors.New("something went wrong 6")),
			result.Error[any](errors.New("something went wrong 7")),
			result.Error[any](errors.New("something went wrong 8")),
			result.Ok[any]("d"),
			result.Error[any](errors.New("something went wrong 9")),
			result.Ok[any]("e"),
			result.Error[any](errors.New("something went wrong 10")),
			result.Error[any](errors.New("something went wrong 11")),
			result.Ok[any]("f"),
			result.Error[any](errors.New("something went wrong 12")),
			result.Error[any](errors.New("something went wrong 13")),
			result.Error[any](errors.New("something went wrong 14")),
		}
	})

	It("should return error values", func() {
		act()

		Expect(observed).To(Equal([]error{
			errors.New("something went wrong 1"),
			errors.New("something went wrong 2"),
			errors.New("something went wrong 3"),
			errors.New("something went wrong 4"),
			errors.New("something went wrong 5"),
			errors.New("something went wrong 6"),
			errors.New("something went wrong 7"),
			errors.New("something went wrong 8"),
			errors.New("something went wrong 9"),
			errors.New("something went wrong 10"),
			errors.New("something went wrong 11"),
			errors.New("something went wrong 12"),
			errors.New("something went wrong 13"),
			errors.New("something went wrong 14"),
		}))
	})
})

var _ = DescribeFunction(result.Combine[any], func() {
	var inputs []result.Result[any]
	var observed result.Result[[]any]

	act := func() {
		observed = result.Combine(inputs...)
	}

	Context("with error results", func() {
		BeforeEach(func() {
			inputs = []result.Result[any]{
				result.Error[any](errors.New("something went wrong 1")),
				result.Error[any](errors.New("something went wrong 2")),
				result.Error[any](errors.New("something went wrong 3")),
				result.Ok[any]("a"),
				result.Error[any](errors.New("something went wrong 4")),
				result.Ok[any]("b"),
				result.Ok[any]("c"),
				result.Error[any](errors.New("something went wrong 5")),
				result.Error[any](errors.New("something went wrong 6")),
				result.Error[any](errors.New("something went wrong 7")),
				result.Error[any](errors.New("something went wrong 8")),
				result.Ok[any]("d"),
				result.Error[any](errors.New("something went wrong 9")),
				result.Ok[any]("e"),
				result.Error[any](errors.New("something went wrong 10")),
				result.Error[any](errors.New("something went wrong 11")),
				result.Ok[any]("f"),
				result.Error[any](errors.New("something went wrong 12")),
				result.Error[any](errors.New("something went wrong 13")),
				result.Error[any](errors.New("something went wrong 14")),
			}
		})

		It("should combine all ok and error values", func() {
			act()

			Expect(observed.IsOk()).To(BeFalse())
			Expect(observed.IsError()).To(BeTrue())
			Expect(observed.OrEmpty()).To(Equal([]any{"a", "b", "c", "d", "e", "f"}))
			Expect(observed.Error()).To(BeAssignableToTypeOf(roar.AggregatedError{}))
			Expect(observed.Error().(roar.AggregatedError).Errors()).To(HaveLen(14))
			Expect(observed.Error().Error()).To(Equal(
				"AggregatedError: multiple errors occurred [" +
					"[0]=something went wrong 1, " +
					"[1]=something went wrong 2, " +
					"[2]=something went wrong 3, " +
					"[3]=something went wrong 4, " +
					"[4]=something went wrong 5, " +
					"[5]=something went wrong 6, " +
					"[6]=something went wrong 7, " +
					"[7]=something went wrong 8, " +
					"[8]=something went wrong 9, " +
					"[9]=something went wrong 10, " +
					"[10]=something went wrong 11, " +
					"[11]=something went wrong 12, " +
					"[12]=something went wrong 13, " +
					"[13]=something went wrong 14]"))
		})
	})

	Context("without error results", func() {
		BeforeEach(func() {
			inputs = []result.Result[any]{
				result.Ok[any]("a"),
				result.Ok[any]("b"),
				result.Ok[any]("c"),
				result.Ok[any]("d"),
				result.Ok[any]("e"),
				result.Ok[any]("f"),
			}
		})

		It("should combine all ok values", func() {
			act()

			Expect(observed.IsOk()).To(BeTrue())
			Expect(observed.IsError()).To(BeFalse())
			Expect(observed.OrEmpty()).To(Equal([]any{"a", "b", "c", "d", "e", "f"}))
			Expect(observed.Error()).To(BeNil())
		})
	})
})
