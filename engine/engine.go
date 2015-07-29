package engine

import (
  "time"
  "fmt"

  "github.com/jonathonharrell/dice"
  "github.com/jonathonharrell/miri-ws-server/websocket"
)

var TheWorld World

func Start() {
  dice.SeedRandom()
  go websocket.StartServer()

  // load in the world, rooms, etc
  TheWorld = NewWorld("The Miri")

  // start the world update loop
  // go RunEvery(WORLD_UPDATE_LOOP_TIMER * time.Second, TheWorld.Update)

  var input string
  fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
