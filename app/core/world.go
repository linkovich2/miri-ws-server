package core

const (
	WorldUpdateLoopTimer = 5
)

type World struct {
	Name         string
	Realms       map[string]*Realm
	sendCallback func(string, string)
}

func (w *World) SetSendCallback(f func(string, string)) {
	w.sendCallback = f
}

func (w *World) Update() {
	for _, r := range w.Realms {
		r.Update(w.sendCallback)
	}
}
