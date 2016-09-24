package dcssa

// Data holds all data for multiple morgues.
type Data struct {
	Runs        []*Run
	FailedReads map[string]string
}

// Run holds data for a specific run.
type Run struct {
	Version string
	Score   int
}
