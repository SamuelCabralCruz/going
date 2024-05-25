//go:build test

package htntp

import (
	"errors"
	. "github.com/SamuelCabralCruz/went/kinggo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httputil"
)

func StubHttpUtilDumpResponse(value string, err error) {
	httpUtilDumpResponse = func(_ *http.Response, _ bool) ([]byte, error) {
		return []byte(value), err
	}
}

var _ = DescribeFunction(DumpResponse, func() {
	var input http.Response
	var observedValue string
	var observedError error

	act := func() { observedValue, observedError = DumpResponse(&input) }

	AfterEach(func() {
		httpUtilDumpResponse = httputil.DumpResponse
	})

	XIt("should dump response", func() {
		// This test sole purpose is to document that format of the returned value.
		// It should not be considered as part of the global test suite.
		url := "https://official-joke-api.appspot.com/random_joke"
		expected := "HTTP/2.0 200 OK Access-Control-Allow-Origin: * Alt-Svc: h3=\":443\"; ma=2592000,h3-29=\":443\"; ma=2592000 Cache-Control: private Content-Type: application/json; charset=utf-8 Date: Mon, 20 May 2024 22:48:39 GMT Etag: W/\"6b-OkN+vnMV/niTobujtDNbrYf1RG0\" Server: Google Frontend Vary: Accept-Encoding X-Cloud-Trace-Context: 152ccb4c0fd5b5639ddfc4451285f576 X-Powered-By: Express  {\"type\":\"general\",\"setup\":\"What is a witch's favorite subject in school?\",\"punchline\":\"Spelling!\",\"id\":247}"
		response, _ := http.Get(url)

		observed, _ := DumpResponse(response)

		Expect(observed).To(Equal(expected))
	})

	Context("with readable response", func() {
		var responseContent string
		var expectedValue string

		BeforeEach(func() {
			responseContent = "some\r\nresponse\r\ncontent"
			expectedValue = "some response content"
			StubHttpUtilDumpResponse(responseContent, nil)
		})

		It("should return sanitized response content", func() {
			act()

			Expect(observedValue).To(Equal(expectedValue))
			Expect(observedError).To(BeNil())
		})
	})

	Context("with error during response reading", func() {
		var readError error

		BeforeEach(func() {
			readError = errors.New("something went wrong")
			StubHttpUtilDumpResponse("", readError)
		})

		It("should return error", func() {
			act()

			Expect(observedValue).To(BeEmpty())
			Expect(observedError).To(Equal(readError))
		})
	})
})
