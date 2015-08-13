package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

var (
	miri  *World
	users map[string]*User
)

func Start() {
	dice.SeedRandom()              // seed rand for dice
	filters.Init()                 // init filter libs (RP filter, profanity filter, language filter, etc)
	users = make(map[string]*User) // init global users map

	ConnectToDatabase("localhost:27017", "miri") //@temp, replace with env vars
	defer CloseDatabaseConnection()

	// auth.CreateUser([]byte("jonathon.harrell@yahoo.com"), []byte("Ex@mple1"))

	StartWebsocketServer()
	RegisterCommandAliases()

	// load in the world, rooms, etc
	miri = &World{"Miri", make(map[string]Realm)}

	// start the world update loop
	go util.RunEvery(WorldUpdateLoopTimer*time.Second, miri.Update)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
