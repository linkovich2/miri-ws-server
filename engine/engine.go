package engine

import (
  "time"
  "fmt"

  "github.com/jonathonharrell/miri/engine/dice"
)

var TheWorld World

func Start() {
  dice.SeedRandom()

  // load in the world, rooms, etc
  TheWorld = NewWorld("The Miri")

  // start the world update loop
  go RunEvery(WORLD_UPDATE_LOOP_TIMER * time.Second, TheWorld.Update)
  //
  var input string
  fmt.Scanln(&input)
}
