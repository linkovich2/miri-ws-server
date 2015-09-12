package engine

const (
	WorldUpdateLoopTimer = 5
)

type World struct {
	Name   string
	Realms map[string]Realm
}

func (w *World) Update() {
	for _, r := range w.Realms {
		r.Update()
	}

	// @todo for testing only, we want to also simulate player actions here
	// to see what the results might be coming back to a client
}
