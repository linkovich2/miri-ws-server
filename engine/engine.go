package engine

import (
	"fmt"
	"time"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"

	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
	"github.com/jonathonharrell/miri-ws-server/engine/routers"
	ws "github.com/jonathonharrell/miri-ws-server/engine/websockets"
	"github.com/jonathonharrell/miri-ws-server/engine/settings/bootstrap"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
	"github.com/jonathonharrell/miri-ws-server/engine/util/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

var (
	miri  *World
	env   settings.Environment
)

func Start() {
	dice.SeedRandom()              // seed rand for dice
	filters.Init()                 // init filter libs (RP filter, profanity filter, language filter, etc)
	settings.LoadEnv()
	env = settings.GetEnv()

	db.ConnectToDatabase(settings.GetEnv().DBHost, settings.GetEnv().DBName)
	defer db.CloseDatabaseConnection()

	bootstrap.Start()

	// @todo temp
	go ws.Hub.Run()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	addr := ":" + strconv.Itoa(env.Port)
	go http.ListenAndServe(addr, n)
	// end temp

	RegisterCommandAliases()

	miri = &World{"Miri", make(map[string]Realm)} // load in the world, rooms, etc

	go util.RunEvery(WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
