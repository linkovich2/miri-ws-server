package core

import (
	"encoding/json"
	"stablelib.com/v1/uniuri"
	"strings"
)

const (
	WorldUpdateLoopTimer = 5
)

type (
	World struct {
		Name         string
		Realms       map[string]*Realm
		sendCallback func(string, string)
	}

	readRoom struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Size        int      `json:"size"`
		Details     []string `json:"details"`
		Entities    []string `json:"entities"`
		Position    Position `json:"position"`
	}
)

func (w *World) SetSendCallback(f func(string, string)) {
	w.sendCallback = f
}

func (w *World) GetSendCallback() func(string, string) {
	return w.sendCallback
}

func (w *World) Update() {
	for _, r := range w.Realms {
		r.Update(w.sendCallback)
	}
}

func GetWorld() *World {
	data := MustAsset("json/world/world.json")
	world := &World{}
	json.Unmarshal(data, world)

	for k, r := range world.Realms {
		data := MustAsset("json/world/" + strings.ToLower(k) + ".json")
		a := map[string]readRoom{}
		err := json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}

		r.Rooms = make(map[string]*Room)

		for positionString, tmp := range a {
			entities := map[string]Entity{}
			for _, k := range tmp.Entities {
				entities[uniuri.New()] = e.get(k)
			}
			// @todo persistance layer probably has something to say about this
			r.Rooms[positionString] = &Room{
				Name:        tmp.Name,
				Description: tmp.Description,
				Size:        tmp.Size,
				Details:     tmp.Details,
				Position:    tmp.Position,
				Entities:    entities,
				Characters:  make(map[string]*Character),
			}
		}
	}

	return world
}
