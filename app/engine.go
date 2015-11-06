package app

import (
	"fmt"

	db "github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/router"
	"github.com/jonathonharrell/miri-ws-server/app/server"
)

var hub *server.ConnectionHub

func Start() {
	g := game.NewGame()
	go g.Start()

	db.ConnectToDatabase(env.DBHost, env.DBName) // create master DB session
	defer db.CloseDatabaseConnection()

	hub = server.GetHub()
	hub.SetHandler(router.NewRouter(g))
	go hub.Run()
	server.Start(env.Port, env.JWTSecretKey)

	var input string
	fmt.Scanln(&input) // we'll probably replace this for non-development environments with something that outputs to a file
}
