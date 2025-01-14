package request

import (
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	host string
}

func NewRequest(host *url.URL) Request {
	return Request{
		host: host.String(),
	}
}

func (r Request) Get() (RequestResult, error) {
	start := time.Now()

	resp, err := http.Get(r.host)
	if err != nil {
		return RequestResult{}, err
	}

	duration := time.Since(start)

	result := NewRequestResult(*resp, start, duration)

	return result, nil
}
