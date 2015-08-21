package engine

import (
	"net/http"
	"strconv"

	"stablelib.com/v1/uniuri"
)

func StartWebsocketServer(port int) {
	go hub.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Error(err.Error())
			return
		}

		c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New()}
		hub.register <- c
		go c.writePump()
		c.readPump()
	})

	addr := ":" + strconv.Itoa(port)
	go http.ListenAndServe(addr, nil)
}
