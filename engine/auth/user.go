package auth

import (
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

const (
	NotAuthenticated = 1
	Authenticated    = 2
	InGame           = 3
)

// this is the main user type, which includes it's websocket connection and other immediate info
type User struct {
	Account    *UserModel
	Connection *websocket.Connection
	IsAdmin    bool
	State      int
}
