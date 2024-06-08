//go:build test

package optional_test

import (
	"github.com/SamuelCabralCruz/went/fn/optional"
	"github.com/SamuelCabralCruz/went/fn/typing"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(optional.Transform[any, any], func() {
	var cut optional.Optional[string]
	var transformer typing.Transformer[string, int]
	var observed optional.Optional[int]

	act := func() {
		observed = optional.Transform(cut, transformer)
	}

	BeforeEach(func() {
		transformer = func(s string) int {
			return len(s)
		}
	})

	Context("with empty optional", func() {
		BeforeEach(func() {
			cut = optional.Empty[string]()
		})

		It("should return empty optional", func() {
			act()

			Expect(observed.IsPresent()).To(BeFalse())
			Expect(observed.IsAbsent()).To(BeTrue())
			Expect(observed.OrEmpty()).To(BeZero())
		})
	})

	Context("with present optional", func() {
		BeforeEach(func() {
			cut = optional.Of("some value")
		})

		It("should return transformed optional", func() {
			act()

			Expect(observed.IsPresent()).To(BeTrue())
			Expect(observed.IsAbsent()).To(BeFalse())
			Expect(observed.OrEmpty()).To(Equal(10))
		})
	})
})

var _ = DescribeFunction(optional.FilterPresent[any], func() {
	var inputs []optional.Optional[any]
	var observed []any

	act := func() {
		observed = optional.FilterPresent(inputs...)
	}

	BeforeEach(func() {
		inputs = []optional.Optional[any]{
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("a"),
			optional.Empty[any](),
			optional.Of[any]("b"),
			optional.Of[any]("c"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("d"),
			optional.Empty[any](),
			optional.Of[any]("e"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("f"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
		}
	})

	It("should return present values", func() {
		act()

		Expect(observed).To(Equal([]any{"a", "b", "c", "d", "e", "f"}))
	})
})

var _ = DescribeFunction(optional.Combine[any], func() {
	var inputs []optional.Optional[any]
	var observed optional.Optional[[]any]

	act := func() {
		observed = optional.Combine(inputs...)
	}

	BeforeEach(func() {
		inputs = []optional.Optional[any]{
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("a"),
			optional.Empty[any](),
			optional.Of[any]("b"),
			optional.Of[any]("c"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("d"),
			optional.Empty[any](),
			optional.Of[any]("e"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Of[any]("f"),
			optional.Empty[any](),
			optional.Empty[any](),
			optional.Empty[any](),
		}
	})

	It("should combine all present values", func() {
		act()

		Expect(observed.IsPresent()).To(BeTrue())
		Expect(observed.IsAbsent()).To(BeFalse())
		Expect(observed.OrEmpty()).To(Equal([]any{"a", "b", "c", "d", "e", "f"}))
	})
})
