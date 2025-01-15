package mock

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

// SetupMockServer setups sample mock responses used by uni tests.
func SetupMockServer() {
	httpmock.Activate()

	httpmock.RegisterResponder(http.MethodGet, "http://example.com/json",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, map[string]string{"test": "test"}))

	httpmock.RegisterResponder(http.MethodGet, "http://example.com/bad-json",
		httpmock.NewStringResponder(http.StatusOK, `{"test": "test"}`))

	httpmock.RegisterResponder(http.MethodGet, "http://example.com/text",
		httpmock.NewStringResponder(http.StatusOK, "test"))

	httpmock.RegisterResponder(http.MethodGet, "http://example.com/empty",
		httpmock.NewStringResponder(http.StatusOK, ""))
}
