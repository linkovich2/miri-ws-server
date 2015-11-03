package engine

// this is just a prelim for dialogue trees, hoping I won't forget about this structure
type DialogueTree struct {
	ID           string
	Text         string
	RaceOption   map[string]bool
	GenderOption map[string]bool
	TraitOption  map[string]bool
	Options      []DialogueTree
}
