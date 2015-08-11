package engine

type room struct {
	id          int
	name        string
	realm       string
	pos         position
	active      bool
	connections []*connection
}

func (r *room) init() {
	r.deactivate()
}

func (r *room) update() {
	// do something, this is the "slow update", world-level update
}

func (r *room) activate() {
	r.active = true
	// here we should activate the "fast updates" for entities in the room
	// spin those off into separate goroutines
}

func (r *room) deactivate() {
	r.active = false

	// we should now "cool down" the "fast updates"
}

func (r *room) broadcast() {

}
