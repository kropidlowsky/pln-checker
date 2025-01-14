package request

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type ResponseValidator struct {
	response http.Response
}

func NewResponseValidator(response http.Response) ResponseValidator {
	return ResponseValidator{
		response: response,
	}
}

// IsJSONContentType checks if the request's Content-Type header is JSON.
func (r ResponseValidator) IsJSONContentType() bool {
	contentType := r.response.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/json")
}

// IsBodyValidJSON validates if response bopy is a valid JSON.
func (r ResponseValidator) IsBodyValidJSON() bool {
	defer r.response.Body.Close()

	bodyBytes, err := io.ReadAll(r.response.Body)
	if err != nil {
		return false
	}

	var body any
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return false
	}

	return true
}
