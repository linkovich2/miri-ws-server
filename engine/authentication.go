package engine

import "encoding/json"

type FormLogin struct {
	Email    string
	Password string
}

func (h *HandlerInterface) CommandNotAuthenticated_AUTHENTICATE(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	hub.Send(form.Email, u.Connection)
	hub.Send("*****", u.Connection)
}

func (h *HandlerInterface) CommandNotAuthenticated_CREATEUSER(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	hub.Send(form.Email, u.Connection)
	hub.Send("*****", u.Connection)
}
