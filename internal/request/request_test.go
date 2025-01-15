package request

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/kropidlowsky/pln-checker/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	tcs := map[string]struct {
		host   string
		result RequestResult
		err    string
	}{
		"correct json response": {
			host: "http://example.com/json",
			result: RequestResult{
				Method:      "GET",
				Host:        "http://example.com/json",
				Status:      http.StatusOK,
				IsJSON:      true,
				IsValidJSON: true,
			},
		},
		"incorrect json response - missing header": {
			host: "http://example.com/bad-json",
			result: RequestResult{
				Method:      "GET",
				Host:        "http://example.com/bad-json",
				Status:      http.StatusOK,
				IsJSON:      false,
				IsValidJSON: true,
			},
		},
		"correct text response": {
			host: "http://example.com/text",
			result: RequestResult{
				Method:      "GET",
				Host:        "http://example.com/text",
				Status:      http.StatusOK,
				IsJSON:      false,
				IsValidJSON: false,
			},
		},
		"invalid request host": {
			host:   "text",
			result: RequestResult{},
			err:    `Get "text": no responder found`,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mock.SetupMockServer()
			defer httpmock.DeactivateAndReset()

			req := NewRequest(tc.host)

			report, err := req.Get()
			if tc.err != "" {
				assert.Error(t, err)
				assert.Zero(t, report)
				return
			}

			assert.NoError(t, err)

			assert.EqualValues(t, tc.result.Method, report.Method)
			assert.EqualValues(t, tc.result.Host, report.Host)
			assert.EqualValues(t, tc.result.Status, report.Status)
			assert.EqualValues(t, tc.result.IsJSON, report.IsJSON)
			assert.EqualValues(t, tc.result.IsValidJSON, report.IsValidJSON)
		})
	}

}
