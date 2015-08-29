package engine

import (
	"encoding/json"
)

type Character struct {
}

func (h *HandlerInterface) CommandAuthenticated_CHARLIST(u *User, args *json.RawMessage) {
	res := &MessageResponse{
		Errors:     nil,
		Success:    true,
		ResponseTo: "charlist",
		Data:       u.Account.Characters,
	}

	hub.Send(res, u.Connection)
}
