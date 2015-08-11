package engine

const (
	worldUpdateLoopTimer = 5
)

type world struct {
	name   string
	realms map[string]realm
}

func (w *world) update() {
	for _, r := range w.realms {
		r.update()
	}

	// @todo for testing only, we want to also simulate player actions here
	// to see what the results might be coming back to a client

	logger.Notice("World update ran")
}
