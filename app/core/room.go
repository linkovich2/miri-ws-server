package core

import "github.com/jonathonharrell/miri-ws-server/app/logger"

const (
	RoomTiny = iota
	RoomSmall
	RoomMedium
	RoomLarge
	RoomHuge
	RoomMassive
)

type Room struct {
	ID          int                   `json:"-"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Detail      string                `json:"detail"`
	Active      bool                  `json:"-"`
	Connections []string              `json:"-"`
	Position    Position              `json:"-"`
	Size        int                   `json:"-"`
	Entities    map[string]Entity     `json:"entities"`
	Characters  map[string]*Character `json:"-"`

	// this is for use with building descriptions or sentences about battles or movement through a room
	// take for instance this slice: []string{"sand", "tents", "noise"}
	// might result in sentences like, "You make your way through the sand|tents|noise" or
	// "The battle rages amidst the sand|tents|noise."
	Details []string `json:"-"`
}

func (r *Room) Update(sendMsg func(string, string)) {
	if r.Active {
		for _, e := range r.Entities {
			e.Update(r, sendMsg)
		}
	}
}

func (r *Room) Add(connectionId string, character *Character) {
	r.Active = true
	r.Connections = append(r.Connections, connectionId)
	r.Characters[connectionId] = character
}

func (r *Room) GetSpeedMod() int {
	// @todo this is work in progress formula
	return r.Size + r.Size/2
}

func (r *Room) Remove(c string) {
	for i, v := range r.Connections {
		if v == c {
			r.Connections = append(r.Connections[:i], r.Connections[i+1:]...)
			delete(r.Characters, c)
		}
	}

	if len(r.Connections) <= 0 {
		r.Active = false
	}
}

func (r *Room) Broadcast(msg string, cb func(string, string)) {
	for _, c := range r.Connections {
		cb(c, msg)
	}
}

func (r *Room) Interact(c *Character, target, action string, cb func(string, string)) {
	if e, exists := r.Entities[target]; exists {
		e.Interact(action, c, r, cb)
	} else {
		logger.Write.Info("Targeted entity [%s] doesn't exist in this room [%s]!", target, r.Position.ToString())
	}
}

func (r *Room) ValidTarget(id string) bool {
	if _, exists := r.Entities[id]; exists {
		return true
	}
	return false
}
