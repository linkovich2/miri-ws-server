package engine

import (
	"flag"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func startWebsocketServer() {
	flag.Parse()
	go hub.run()
	http.HandleFunc("/", serveWs)
	go http.ListenAndServe(*addr, nil)
}
