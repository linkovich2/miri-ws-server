package engine

import "encoding/json"

type (
	FormLogin struct {
		Email    string
		Password string
	}

	LoginResponse struct { // @todo create a more generalized response struct
		Errors     []string `json:"errors"`
		Success    bool     `json:"success"`
		ResponseTo string   `json:"response_to"`
	}

	SignupResponse struct {
		Errors     []string `json:"errors"`
		Success    bool     `json:"success"`
		ResponseTo string   `json:"response_to"`
	}
)

func (h *HandlerInterface) CommandNotAuthenticated_AUTHENTICATE(u *User, args *json.RawMessage) {
	form := &FormLogin{}
	err := json.Unmarshal(*args, &form)

	if err != nil {
		// @todo something is probably missing from the JSON
		return
	}

	// hub.Send(form.Email, u.Connection)
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
	}

	res := &SignupResponse{errors, success, "createuser"}
	j, _ := json.Marshal(res)
	hub.Send(j, u.Connection)
}
