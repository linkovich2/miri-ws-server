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

	success, errors := Authenticate(form.Email, form.Password)
	if success {
		u.State = Authenticated
		logger.Info("User logged in: %s", form.Email)
	}

	hub.Send(&MessageResponse{Errors: errors, Success: success, ResponseTo: "authenticate"}, u.Connection)
}

func (h *HandlerInterface) CommandNotAuthenticated_CREATEUSER(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	errors := CreateUser(form.Email, form.Password)
	success := len(errors) <= 0
	if success {
		u.State = Authenticated
		logger.Info("New User: %s", form.Email)
	}

	hub.Send(&MessageResponse{Errors: errors, Success: success, ResponseTo: "createuser"}, u.Connection)
}
