package engine

import (
	"fmt"
	"time"

	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
	"github.com/jonathonharrell/miri-ws-server/engine/settings/bootstrap"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

var (
	miri  *World
	users map[string]*User
	env   settings.Environment
)

func Start() {
	dice.SeedRandom()              // seed rand for dice
	filters.Init()                 // init filter libs (RP filter, profanity filter, language filter, etc)
	users = make(map[string]*User) // init global users map
	settings.LoadEnv()
	env = settings.GetEnv()

	db.ConnectToDatabase(settings.GetEnv().DBHost, settings.GetEnv().DBName)
	defer db.CloseDatabaseConnection()

	bootstrap.Init()

	StartWebsocketServer(env.Port)
	RegisterCommandAliases()

	miri = &World{"Miri", make(map[string]Realm)} // load in the world, rooms, etc

	go util.RunEvery(WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
