package request

import (
	"io"
	"net/http"
	"strings"
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

func TestIsBodyValidJSON(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		body   string
		isJSON bool
	}{
		"correct response having JSON body": {
			body: `{
			"test": "test"
			}`,
			isJSON: true,
		},
		"correct response having array JSON body": {
			body: `[{
			"test": "test"
			}]`,
			isJSON: true,
		},
		"incorrect response having XML body": {
			body: `<?xml version="1.0" encoding="UTF-8" ?>
			<root>
				<test>test</test>
			</root>`,
			isJSON: false,
		},
		"incorrect response having text body": {
			body:   `text`,
			isJSON: false,
		},
	}

	for name, tc := range tcs {
		tc := tc

		t.Run(name, func(t *testing.T) {
			bodyReader := strings.NewReader(tc.body)
			bodyCloser := io.NopCloser(bodyReader)
			defer bodyCloser.Close()

			response := http.Response{Body: bodyCloser}
			responseValidator := NewResponseValidator(response)

			isJSON := responseValidator.IsBodyValidJSON()
			assert.EqualValues(t, tc.isJSON, isJSON)
		})

	}

}
