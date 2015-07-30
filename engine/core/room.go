package core

import (

)

type Room struct {
  ID       int
  Name     string
  Realm    string
  Position Position
  Active   bool
}

func (room *Room) Init() {
  room.Deactivate()
}

func (room *Room) Update() {
  // do something, this is the "slow update", world-level update
}

func (room *Room) Activate() {
  room.Active = true
  // here we should activate the "fast updates" for entities in the room
  // spin those off into separate goroutines
}

func (room *Room) Deactivate() {
  room.Active = false

  // we should now "cool down" the "fast updates"
}
