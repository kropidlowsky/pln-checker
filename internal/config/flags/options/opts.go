package options

import (
	"flag"
	"net/url"
)

// Opts represents common options.
type Opts struct {
	Host host
}

type host struct {
	host *url.URL
}

var _ flag.Value = (*host)(nil)

func (h *host) Set(host string) error {
	u, err := url.ParseRequestURI(host)
	if err != nil {
		return err
	}

	h.host = u

	return nil
}

func (h host) String() string {
	if h.host == nil {
		return ""
	}

	return h.host.String()
}
