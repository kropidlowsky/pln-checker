package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsJSONContentType(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		header            http.Header
		isJSONContentType bool
	}{
		"response having correct JSON Content-Type header": {
			header:            http.Header{"Content-Type": []string{"application/json; charset=utf-8"}},
			isJSONContentType: true,
		},
		"response having short correct JSON Content-Type header": {
			header:            http.Header{"Content-Type": []string{"application/json"}},
			isJSONContentType: true,
		},
		"response having incorrect JSON Content-Type header": {
			header:            http.Header{"Content-Type": []string{"json"}},
			isJSONContentType: false,
		},
		"response having short incorrect Content-Type header": {
			header:            http.Header{"Content-Type": []string{"application/xml"}},
			isJSONContentType: false,
		},
	}

	for name, tc := range tcs {
		tc := tc
		t.Run(name, func(t *testing.T) {
			response := http.Response{Header: tc.header}
			responseValidator := NewResponseValidator(response)
			isJSONContentType := responseValidator.IsJSONContentType()
			assert.EqualValues(t, tc.isJSONContentType, isJSONContentType)
		})

	}

}
