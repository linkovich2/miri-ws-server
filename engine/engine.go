package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

var (
	miri world
	users map[string]*user
)

func Start() {
	dice.SeedRandom()              // seed rand for dice
	filters.Init()                 // init filter libs (RP filter, profanity filter, language filter, etc)
	users = make(map[string]*user) // init global users map

	connectToDatabase("localhost:27017", "miri") //@temp, replace with env vars
	defer closeDatabaseConnection()

	// auth.CreateUser([]byte("jonathon.harrell@yahoo.com"), []byte("Ex@mple1"))

	startWebsocketServer()
	attachMessageHandlers() // probably don't need this @todo
	registerCommandAliases()

	// load in the world, rooms, etc
	miri = world{"Miri", make(map[string]realm)}

	// start the world update loop
	go util.RunEvery(worldUpdateLoopTimer*time.Second, miri.update)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
