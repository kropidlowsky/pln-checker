package options

// LoadOpts represents load testing options.
type LoadOpts struct {
	Opts

	Rate      uint
	Frequency uint
}
