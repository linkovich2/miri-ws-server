package engine

import "encoding/json"

type formLogin struct {
	email    string
	password string
}

func cmdAuthenticate(u *user, args *json.RawMessage) {
	form := &formLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// something is probably missing from the JSON
		return
	}

	hub.Send([]byte("Trying to authenticate"), u.connection)
}
