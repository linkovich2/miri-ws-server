package core

type Room struct {
	ID          int      `json:"-"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Active      bool     `json:"-"`
	Connections []string `json:"-"`
	Position    Position
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
