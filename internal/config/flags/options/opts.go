package options

import "net/url"

// Opts represents common options.
type Opts struct {
	Host host
}

type host struct {
	host *url.URL
}

func (h *host) Set(host string) error {
	u, err := url.Parse(host)
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
