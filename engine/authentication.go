package engine

import "encoding/json"

func cmdAuthenticate(u *user, args *json.RawMessage) {
	form := &formLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// something is probably missing from the JSON
		return
	}

	hub.Send([]byte("Trying to authenticate"), u.connection)
}
