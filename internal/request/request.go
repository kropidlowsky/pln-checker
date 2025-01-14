package request

import (
	"net/http"
	"time"
)

type Request struct {
	client *http.Client
	host   string
}

func NewRequest(host string) Request {
	return Request{
		client: http.DefaultClient,
		host:   host,
	}
}

func (r Request) Get() (RequestResult, error) {
	start := time.Now()

	req, err := r.request()
	if err != nil {
		return RequestResult{}, err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return RequestResult{}, err
	}

	defer resp.Body.Close()

	duration := time.Since(start)

	result := NewRequestResult(*resp, start, duration)

	return result, nil
}

// request creates the request which includes User-Agent header required by some APIs.
func (r Request) request() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, r.host, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "")

	return req, nil
}
