package core

const (
	RoomTiny = iota
	RoomSmall
	RoomMedium
	RoomLarge
	RoomHuge
	RoomMassive
)

type Room struct {
	ID          int      `json:"-"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Detail      string   `json:"detail"`
	Active      bool     `json:"-"`
	Connections []string `json:"-"`
	Position    Position `json:"-"`
	Size        int      `json:"-"`

	// this is for use with building descriptions or sentences about battles or movement through a room
	// take for instance this slice: []string{"sand", "tents", "noise"}
	// might result in sentences like, "You make your way through the sand|tents|noise" or
	// "The battle rages amidst the sand|tents|noise."
	Details []string `json:"-"`
}

func (r *Room) Update() {
	if r.Active {
		// do something, this is the "slow update", world-level update
	}
}

func (r *Room) Add(c string) {
	r.Active = true
	r.Connections = append(r.Connections, c)
}

func (r *Room) GetSpeedMod() int {
	// @todo this is work in progress formula
	return r.Size + r.Size/2
}

func (r *Room) Remove(c string) {
	for i, v := range r.Connections {
		if v == c {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
		}
	}

	if len(r.Connections) <= 0 {
		r.Active = false
	}
}
