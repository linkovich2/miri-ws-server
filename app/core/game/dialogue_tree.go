package game

// just a prelim for dialogue trees
type DialogueTree struct {
	ID           string
	Text         string
	RaceOption   map[string]bool
	GenderOption map[string]bool
	TraitOption  map[string]bool
	Options      []DialogueTree
}
