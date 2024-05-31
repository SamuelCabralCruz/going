//go:build test

package gomicron_test

import (
	. "github.com/SamuelCabralCruz/went/gomicron"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("matcher", func() {
	It("should report properly", func() {
		Expect("true").To(BeAStringOfLength(6))
		//Expect(true).To(BeAStringOfLength(6))
	})
})
