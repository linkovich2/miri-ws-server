package controllers

import (
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	ws "github.com/jonathonharrell/miri-ws-server/engine/websockets"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, logger.Write.Error("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return env.SessionKey, nil
		}
	})

	if err == nil && token.Valid) {
		userId, err := getUserId(r)
		if err != nil {
			logger.Write.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ws.WebsocketServe(w, r, userId)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
