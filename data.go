package dcssa

// Data holds all data for multiple morgues.
type Data struct {
	Runs        []*Run
	FailedReads map[string]string
	Stats       Stats
}

// NewData creates a new Data struct and sets some default values.
func NewData() *Data {
	return &Data{
		Runs:        make([]*Run, 0),
		FailedReads: make(map[string]string),
	}
}

// Stats represent stats calculated from multiple runs.
type Stats struct {
	Runs           int
	Wins           int
	RacesWon       []string
	BackgroundsWon []string
	CharactersWon  []string
}

// NewStats creates a new Stats struct with defaults.
func NewStats() Stats {
	return Stats{
		RacesWon:       make([]string, 0),
		BackgroundsWon: make([]string, 0),
		CharactersWon:  make([]string, 0),
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
	Attributes  map[string][]string
	Escaped     bool
	Skills      []Skill
	Spells      []Spell
	Mutations   []string
}

// Skill represents a skill and its level.
type Skill struct {
	Name         string
	Level        int
	LevelDecimal int
	State        SkillState
}

// SkillState represents all states a skill can be in.
type SkillState string

// All possible SkillStates
const (
	UNUSED   = SkillState("unused")
	USING    = SkillState("using")
	TRAINING = SkillState("training")
	FOCUSING = SkillState("focusing")
	MASTERED = SkillState("mastered")
)

// Spell represent a known spell.
type Spell struct {
	Name    string
	Type    string
	Power   string
	Failure int
	Level   int
	Hunger  string
}

// NewRun creates a new Run without nils.
func NewRun() *Run {
	return &Run{
		Stats:       make(map[string]string),
		Resistances: make(map[string]string),
		Equipped:    make([]string, 0),
		Attributes:  make(map[string][]string),
		Skills:      make([]Skill, 0),
		Spells:      make([]Spell, 0),
		Mutations:   make([]string, 0),
	}
}
