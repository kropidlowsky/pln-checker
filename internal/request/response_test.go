package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsJSONContentType(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		response          http.Response
		isJSONContentType bool
	}{
		"response having correct JSON Content-Type header": {
			response:          http.Response{Header: http.Header{"Content-Type": []string{"application/json; charset=utf-8"}}},
			isJSONContentType: true,
		},
		"response having short correct JSON Content-Type header": {
			response:          http.Response{Header: http.Header{"Content-Type": []string{"application/json"}}},
			isJSONContentType: true,
		},
		"response having incorrect JSON Content-Type header": {
			response:          http.Response{Header: http.Header{"Content-Type": []string{"son"}}},
			isJSONContentType: false,
		},
		"response having short incorrect Content-Type header": {
			response:          http.Response{Header: http.Header{"Content-Type": []string{"application/xml"}}},
			isJSONContentType: false,
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			responseValidator := NewResponseValidator(tc.response)
			isJSONContentType := responseValidator.IsJSONContentType()
			assert.EqualValues(t, tc.isJSONContentType, isJSONContentType)
		})

	}

}
