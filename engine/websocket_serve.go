package engine

import (
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"stablelib.com/v1/uniuri"

	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/routers"
)

func StartWebsocketServer(port int) {
	go hub.run()

	router := routers.InitRoutes()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Write.Error(err.Error())
			return
		}

		// @todo check bearer token before making the connection, or else send a 401

		c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New()}
		hub.register <- c
		go c.writePump()
		c.readPump()
	}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(router)

	addr := ":" + strconv.Itoa(port)
	go http.ListenAndServe(addr, n)
}
