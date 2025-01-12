package options

import "net/url"

// Opts represents common options.
type Opts struct {
	Host *url.URL
}

func (o *Opts) Set(host string) error {
	u, err := url.Parse(host)
	if err != nil {
		return err
	}

	o.Host = u

	return nil
}
