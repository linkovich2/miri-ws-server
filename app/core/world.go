package core

const (
	WorldUpdateLoopTimer = 5
)

type World struct {
	Name   string
	Realms map[string]*Realm
}

func (w *World) Update() {
	for _, r := range w.Realms {
		r.Update()
	}
}
