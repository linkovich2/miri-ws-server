package engine

type Room struct {
	ID          int
	Name        string
	Realm       string
	Position    Position
	Active      bool
	Description string
	Connections []*Connection
}

func (r *Room) Init() {
	r.Deactivate()
}

func (r *Room) Update() {
	// do something, this is the "slow update", world-level update
}

func (r *Room) Activate() {
	r.Active = true
	// here we should activate the "fast updates" for entities in the room
	// spin those off into separate goroutines
}

func (r *Room) Deactivate() {
	r.Active = false

	// we should now "cool down" the "fast updates"
}

func (r *Room) Broadcast() {
}
