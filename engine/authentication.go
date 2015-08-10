package engine

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

func cmdAuthenticate(u *auth.User, args *json.RawMessage) {
	form := &auth.LoginData{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// something is probably missing from the JSON
		return
	}

	hub.Send([]byte("Trying to authenticate"), u.Connection)
}

func onConnectCallback(c *websocket.Connection) {
	users[c.ID] = &auth.User{Connection: c, State: auth.NotAuthenticated}
	// maybe we should also try to authenticate, if we want to use cookies or whatever
}
