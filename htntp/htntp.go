package htntp

import (
	"github.com/SamuelCabralCruz/went/fn/result"
	"net/http"
	"net/http/httputil"
	"strings"
)

var httpUtilDumpResponse = httputil.DumpResponse

func DumpResponse(response *http.Response) (string, error) {
	return result.Transform(result.FromAssertion(httpUtilDumpResponse(response, true)), func(value []byte) string {
		return string(value)
	}).Map(func(value string) string {
		return strings.ReplaceAll(value, "\r\n", " ")
	}).Get()
}
