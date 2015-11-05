package controllers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	ws "github.com/jonathonharrell/miri-ws-server/engine/websockets"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	userId, err := getUserId(r)
	if err != nil {
		logger.Write.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ws.WebsocketServe(w, r, userId)
}
