package engine

const (
	WorldUpdateLoopTimer = 5
)

type World struct {
	Name   string
	Realms map[string]Realm
}

func (world *World) Update() {
	for _, r := range world.Realms {
		r.Update()
	}

	// @todo for testing only, we want to also simulate player actions here
	// to see what the results might be coming back to a client

	logger.Notice("World update ran")
}

func NewWorld(name string) World {
	world := World{name, make(map[string]Realm, 8)}
	return world
}
