package server

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"

	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/routers"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
	"github.com/jonathonharrell/miri-ws-server/engine/settings/bootstrap"
	ws "github.com/jonathonharrell/miri-ws-server/engine/websockets"
)

func Start() {
	settings.LoadEnv()
	env := settings.GetEnv()

	db.ConnectToDatabase(settings.GetEnv().DBHost, settings.GetEnv().DBName)
	defer db.CloseDatabaseConnection()

	bootstrap.Start()

	go ws.Hub.Run()

	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	addr := ":" + strconv.Itoa(env.Port)
	go http.ListenAndServe(addr, n)
}
