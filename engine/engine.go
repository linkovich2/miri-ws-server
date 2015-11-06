package engine

import (
	"fmt"
	"time"

	db "github.com/jonathonharrell/miri-ws-server/engine/database"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

func Start() {
	dice.SeedRandom() // seed rand for dice
	filters.Init()    // init filter libs (RP filter, profanity filter, language filter, etc)

	miri := &World{"Miri", make(map[string]Realm)}                  // load in the world, rooms, etc
	go util.RunEvery(WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	db.ConnectToDatabase(env.DBHost, env.DBName) // create master DB session
	defer db.CloseDatabaseConnection()

	go Hub.Run()
	serve()

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
