package engine

import (
	"flag"
	"net/http"

	"stablelib.com/v1/uniuri"
)

var addr = flag.String("addr", ":8080", "http service address")

func StartWebsocketServer() {
	flag.Parse()
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

	go http.ListenAndServe(*addr, nil)
}
