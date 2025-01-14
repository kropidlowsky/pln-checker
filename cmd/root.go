package cmd

import (
	"fmt"
	"os"

	"github.com/kropidlowsky/pln-checker/internal/attacker"
	"github.com/kropidlowsky/pln-checker/internal/config/options"
	"github.com/spf13/cobra"
)

var opts = &options.LoadOpts{}

var rootCmd = &cobra.Command{
	Use:   "pln-checker",
	Short: "load test",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	attacker := attacker.NewAttacker(*opts)
	attacker.InfiniteAttack()
}

func init() {
	rootCmd.PersistentFlags().Var(&opts.Host, "Host", "Host to send the request to")
	rootCmd.PersistentFlags().UintVar(&opts.Rate, "X", 0, "Numer of requests per the frequency (Y)")
	rootCmd.PersistentFlags().UintVar(&opts.Frequency, "Y", 0, "Interval in seconds for making X requests at once")
}
