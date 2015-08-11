package engine

import "encoding/json"

func cmdAuthenticate(u *User, args *json.RawMessage) {
	form := &LoginData{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// something is probably missing from the JSON
		return
	}

	hub.Send([]byte("Trying to authenticate"), u.Connection)
}

func onConnectCallback(c *Connection) {
	users[c.ID] = &User{Connection: c, State: NotAuthenticated}
	// maybe we should also try to authenticate, if we want to use cookies or whatever
}
