package request

import (
	"net/http"
	"time"
)

type RequestResult struct {
	Method      string
	Host        string
	Start       string
	Duration    string
	Status      int
	IsJSON      bool
	IsValidJSON bool
}

func NewRequestResult(resp http.Response, start time.Time, duration time.Duration) RequestResult {
	responseValidator := NewResponseValidator(resp)

	return RequestResult{
		Method:      resp.Request.Method,
		Host:        resp.Request.URL.String(),
		Start:       start.String(),
		Duration:    duration.String(),
		Status:      resp.StatusCode,
		IsJSON:      responseValidator.IsJSONContentType(),
		IsValidJSON: responseValidator.IsBodyValidJSON(),
	}
}
