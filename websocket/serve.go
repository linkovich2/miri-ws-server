package websocket

import (
	"flag"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func StartServer() *Hub {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", ServeWS)
	go http.ListenAndServe(*addr, nil)
	return &h
}
