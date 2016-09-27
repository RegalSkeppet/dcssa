package dcssa

// Data holds all data for multiple morgues.
type Data struct {
	Runs        []*Run
	FailedReads map[string]string
}

// NewData creates a new Data struct and sets some default values.
func NewData() *Data {
	return &Data{
		Runs:        make([]*Run, 0),
		FailedReads: make(map[string]string),
	}
}

// Run holds data for a specific run.
type Run struct {
	Version     string
	Score       int
	Name        string
	Title       string
	Race        string
	Background  string
	Turns       int
	Time        string
	Stats       map[string]string
	Resistances map[string]string
	Equipped    []string
	Escaped     bool
	Orb         bool
	Runes       []string
}

// NewRun creates a new Run without nils.
func NewRun() *Run {
	return &Run{
		Stats:       make(map[string]string),
		Resistances: make(map[string]string),
		Equipped:    make([]string, 0),
		Runes:       make([]string, 0),
	}
}
