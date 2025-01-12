package flags

// StressFlags represents stress testing options.
type StressFlags struct {
	Flags

	Rate      uint
	Frequency uint
}
