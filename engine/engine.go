package engine

import (
  // "time"
  "fmt"

  "github.com/jonathonharrell/dice"
  "github.com/jonathonharrell/miri-ws-server/engine/websocket"
  "github.com/jonathonharrell/miri-ws-server/engine/core"
  // "github.com/jonathonharrell/miri-ws-server/engine/util"
)

var TheWorld core.World

// @temp
type User struct {
  Connection *websocket.Connection
}

var users []*User

func Start() {
  dice.SeedRandom()
  hub := websocket.StartServer()
  hub.SetOnConnectCallback(func(c *websocket.Connection) {
    // here we should probably give the connection an id for reference and association
    // and attach that ID to the connection itself
    // This way we can reference Users with a connection ID
  })

  hub.SetOnMessageCallback(func(m *websocket.Message) {
    // since we'll have an ID value on the connection, we can reference our list of users
    hub.Send(m.Payload, m.Connection) // pong back the message for now
  })

  // load in the world, rooms, etc
  TheWorld = core.NewWorld("The Miri")

  // start the world update loop
  // go util.RunEvery(WORLD_UPDATE_LOOP_TIMER * time.Second, TheWorld.Update)

  var input string
  fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
