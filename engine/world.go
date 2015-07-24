package engine

import (

)

const (
  WORLD_UPDATE_LOOP_TIMER = 1 // @temp
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

  log.Notice("World update ran")
}


func NewWorld(name string) World {
  world := World{name, make(map[string]Realm, 8)}
  return world
}
