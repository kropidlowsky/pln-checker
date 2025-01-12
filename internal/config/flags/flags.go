package flags

import "net/url"

// Flags represents common flags.
type Flags struct {
	Host *url.URL
}
