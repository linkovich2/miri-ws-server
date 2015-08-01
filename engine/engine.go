package engine

import (
  "time"
  "fmt"

  "gopkg.in/mgo.v2"

  "github.com/jonathonharrell/dice"
  "github.com/jonathonharrell/miri-ws-server/engine/websocket"
  "github.com/jonathonharrell/miri-ws-server/engine/core"
  "github.com/jonathonharrell/miri-ws-server/engine/util"
  "github.com/jonathonharrell/miri-ws-server/engine/database"
  "github.com/jonathonharrell/miri-ws-server/engine/auth"
)

var (
  world core.World
  db *mgo.Database
  users []*auth.User
)

func Start() {
  dice.SeedRandom()

  db = database.Connect("localhost:27017", "miri") //@temp, replace with env vars

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
  world = core.NewWorld("The Miri")

  // start the world update loop
  go util.RunEvery(core.WorldUpdateLoopTimer * time.Second, world.Update)

  var input string
  fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
