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

func (r *Room) Add(c string) {
	r.Active = true
	r.Connections = append(r.Connections, c)
}

func (r *Room) Remove(c string) {
	if len(r.Connections) <= 0 {
		r.Active = false
	}

	for i, v := range r.Connections {
		if v == c {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
		}
	}
}
