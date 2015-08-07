package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/core"
	"github.com/jonathonharrell/miri-ws-server/engine/database"
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
	dice.SeedRandom()                   // seed rand for dice
	filters.Init()                      // init filter libs (RP filter, profanity filter, language filter, etc)
	users = make(map[string]*auth.User) // init global users map

	database.Connect("localhost:27017", "miri") //@temp, replace with env vars
	defer database.Close()                      // when the program exits, close the mongo connection

	// auth.CreateUser([]byte("jonathon.harrell@yahoo.com"), []byte("Ex@mple1"))

	hub = websocket.StartServer()

	attachMessageHandlers()

	// load in the world, rooms, etc
	world = core.NewWorld("The Miri")

	// start the world update loop
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, world.Update)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
