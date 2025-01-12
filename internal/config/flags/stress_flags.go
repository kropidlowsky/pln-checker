package flags

import (
	"flag"

	"github.com/kropidlowsky/pln-checker/internal/config/flags/options"
)

func ParseStressFlags(args []string) (*options.StressOpts, error) {
	fs := flag.NewFlagSet("stress tests", flag.ExitOnError)

	opts := &options.StressOpts{}

	fs.Var(&opts.Host, "Host", "Host to send the request to")
	fs.UintVar(&opts.Rate, "X", 0, "Numer of requests per the frequency (Y)")
	fs.UintVar(&opts.Frequency, "Y", 0, "Interval in seconds for making X requests at once")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return opts, nil
}
