package core

import (
	"github.com/jonathonharrell/miri-ws-server/app/logger"

	"errors"
	"strings"
)

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

// send a message to all active connections
func (r *Room) Broadcast(msg string, cb func(string, string)) {
	for _, c := range r.Connections {
		cb(c, msg)
	}
}

// send a message to all but the specified character
func (r *Room) BroadcastToAllButCharacter(msg string, c *Character, cb func(string, string)) {
	for id, character := range r.Characters {
		if character != c {
			cb(id, msg)
		}
	}
}

// send a message to the specified character
func (r *Room) Message(msg string, c *Character, cb func(string, string)) {
	for id, character := range r.Characters {
		if character == c {
			cb(id, msg)
		}
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
	// @todo it should also be a valid target if it is a character
	if _, exists := r.Entities[id]; exists {
		return true
	}
	return false
}

func (r *Room) GetTarget(id string) (string, error) {
	// @todo this should check characters list as well, and return a "targetting" name
	if !r.ValidTarget(id) {
		return "", errors.New("Requested entity not available in room.")
	}

	e := r.Entities[id].(*ComponentBag)
	descriptor := e.Properties.ValueOf("descriptor")
	if descriptor == "" {
		return strings.ToLower(e.Name), nil
	}

	return descriptor, nil
}
