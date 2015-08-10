package auth

import (
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

const (
	NotAuthenticated = iota
	Authenticated
	InGame
)

type (
	LoginData struct {
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
