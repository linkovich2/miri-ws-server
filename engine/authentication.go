package engine

import "encoding/json"

type FormLogin struct {
	Email    string
	Password string
}

type AuthenticateResponse struct {
	IsAdmin bool `json:"is_admin"`
}

func (h *HandlerInterface) CommandNotAuthenticated_AUTHENTICATE(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	success, errors := Authenticate(form.Email, form.Password, u)

	res := &AuthenticateResponse{}
	if success {
		res.IsAdmin = u.Account.IsAdmin
	}

	hub.Send(&MessageResponse{Errors: errors, Success: success, ResponseTo: "authenticate", Data: res}, u.Connection)
}

func (h *HandlerInterface) CommandNotAuthenticated_CREATEUSER(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	errors := CreateUser(form.Email, form.Password, u)
	success := len(errors) <= 0

	res := &AuthenticateResponse{}
	if success {
		res.IsAdmin = u.Account.IsAdmin
	}

	hub.Send(&MessageResponse{Errors: errors, Success: success, ResponseTo: "createuser", Data: res}, u.Connection)
}
