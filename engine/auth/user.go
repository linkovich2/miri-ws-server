package auth

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

var hub *websocket.Hub

const (
	NotAuthenticated = 1
	Authenticated    = 2
	InGame           = 3
)

type (
	loginData struct {
		Email    string
		Password string
	}

	User struct {
		Account    *UserModel
		Connection *websocket.Connection
		IsAdmin    bool
		State      int
	}
)

func SetHub(h *websocket.Hub) {
	hub = h
}

func CmdAuthenticate(u *User, args *json.RawMessage) {
	form := &loginData{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// something is probably missing from the JSON
		return
	}

	hub.Send([]byte("Trying to authenticate"), u.Connection)
}
