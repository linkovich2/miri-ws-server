package app

import (
	"fmt"
	"time"

	db "github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/server"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/util/dice"
	"github.com/jonathonharrell/miri-ws-server/app/util/filters"
)

var hub server.ConnectionHub

func Start() {
	dice.SeedRandom() // seed rand for dice
	filters.Init()    // init filter libs (RP filter, profanity filter, language filter, etc)

	miri := &core.World{"Miri", make(map[string]core.Realm)}                  // load in the world, rooms, etc
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	db.ConnectToDatabase(env.DBHost, env.DBName) // create master DB session
	defer db.CloseDatabaseConnection()

	hub = server.GetHub()
	go hub.Run()
	server.Start(env.Port, env.JWTSecretKey)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}

func InitGameData() {
	// @todo stub : this should just gather up everything in the "content" package into suitable globals, so that they are always available
}
