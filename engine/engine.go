package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/core"
	"github.com/jonathonharrell/miri-ws-server/engine/database"
	"github.com/jonathonharrell/miri-ws-server/engine/message_handler"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

var (
	world core.World
	users map[string]*auth.User
	hub   *websocket.Hub
)

func Start() {
	dice.SeedRandom()
	filters.Init()
	users = make(map[string]*auth.User) // init global users map

	database.Connect("localhost:27017", "miri") //@temp, replace with env vars
	defer database.Close()                      // when the program exits, close the mongo connection

	// auth.CreateUser([]byte("jonathon.harrell@yahoo.com"), []byte("Ex@mple1"))

	hub = websocket.StartServer()
	hub.SetOnConnectCallback(func(c *websocket.Connection) {
		users[c.ID] = &auth.User{Connection: c, State: auth.NotAuthenticated}
		// maybe we should also try to authenticate, if we want to use cookies or whatever
	})

	hub.SetOnMessageCallback(func(m *websocket.Message) {
		message_handler.Interpret(m, users[m.Connection.ID])
	})

	message_handler.Init() // set up message handler and router
	// import handlers
	message_handler.AddHandler(auth.NotAuthenticated, "say", CmdSay) // @todo move this to a message_handler_register file

	// create error handlers for message handler
	message_handler.InvalidStateHandler = func(u *auth.User, args ...interface{}) {
		hub.Send([]byte("State not valid for some reason."), u.Connection)
	}
	message_handler.InvalidHandlerIndex = func(u *auth.User, args ...interface{}) {
		hub.Send([]byte("Command not found."), u.Connection)
	}

	// load in the world, rooms, etc
	world = core.NewWorld("The Miri")

	// start the world update loop
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, world.Update)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
