package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

var (
	world World
	users map[string]*User
)

func Start() {
	dice.SeedRandom()              // seed rand for dice
	filters.Init()                 // init filter libs (RP filter, profanity filter, language filter, etc)
	users = make(map[string]*User) // init global users map

	connectToDatabase("localhost:27017", "miri") //@temp, replace with env vars
	defer closeDatabaseConnection()

	// auth.CreateUser([]byte("jonathon.harrell@yahoo.com"), []byte("Ex@mple1"))

	startWebsocketServer()
	attachMessageHandlers()

	// load in the world, rooms, etc
	world = NewWorld("The Miri")

	// start the world update loop
	go util.RunEvery(WorldUpdateLoopTimer*time.Second, world.Update)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
