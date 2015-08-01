package auth

import (
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

// this is the main user type, which includes it's websocket connection and other immediate info
type User struct {
	Account *UserModel
	Connection *websocket.Connection
	IsAdmin bool
	IsAuthenticated bool
}
