package flags

import (
	"flag"

	"github.com/kropidlowsky/pln-checker/internal/config/flags/options"
)

// StressFlags represents stress testing options.
type StressFlags struct {
	Flags

	Rate      uint
	Frequency uint
}

func ParseStressFlags() {
	fs := flag.NewFlagSet("stress tests", flag.ExitOnError)

	opts := &options.StressOpts{}

	fs.Var(&opts.Host, "host", "")
}
