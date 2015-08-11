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
