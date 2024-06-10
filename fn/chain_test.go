//go:build test

package fn_test

import (
	"github.com/SamuelCabralCruz/went/fn"
	"github.com/SamuelCabralCruz/went/fn/typing"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeFunction(fn.ChainMappers[any], func() {
	var mappers []typing.Mapper[string]
	var observed typing.Mapper[string]

	act := func() {
		observed = fn.ChainMappers(mappers...)
	}

	BeforeEach(func() {
		mappers = []typing.Mapper[string]{
			func(s string) string { return s + " - " },
			func(s string) string { return s + "m" },
			func(s string) string { return s + "a" },
			func(s string) string { return s + "p" },
			func(s string) string { return s + "p" },
			func(s string) string { return s + "e" },
			func(s string) string { return s + "d" },
		}
	})

	It("should chain mappers sequentially", func() {
		act()

		Expect(observed("some input value")).To(Equal("some input value - mapped"))
	})
})

var _ = DescribeFunction(fn.ChainTransformers[any, any, any], func() {
	var transformer1 typing.Transformer[string, []byte]
	var transformer2 typing.Transformer[[]byte, int]
	var observed typing.Transformer[string, int]

	act := func() {
		observed = fn.ChainTransformers(transformer1, transformer2)
	}

	BeforeEach(func() {
		transformer1 = func(v string) []byte {
			return []byte(v)
		}
		transformer2 = func(v []byte) int {
			return len(v)
		}
	})

	It("should chain transformers sequentially", func() {
		act()

		Expect(observed("some input value")).To(Equal(16))
	})
})
