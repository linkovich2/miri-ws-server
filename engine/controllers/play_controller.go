package controllers

import (
	ws "github.com/jonathonharrell/miri-ws-server/engine/websockets"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ws.WebsocketServe(w, r)
}
