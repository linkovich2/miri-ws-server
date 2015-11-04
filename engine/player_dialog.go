package engine

import (
	// "encoding/json"
	//
	// "github.com/jonathonharrell/miri-ws-server/engine/logger"
	// "github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

type ChatArgs struct {
	Input string
}

// func (h *HandlerInterface) Command_SAY(u *User, args *json.RawMessage) {
// 	logger.Write.Info("Called 'say'!")
// 	chat := &ChatArgs{}
// 	json.Unmarshal(*args, &chat)
//
// 	s := filters.ReplaceProfanity(chat.Input)
// 	// @todo RP filter and any other filters we want chat to go through
//
// 	hub.Send(&MessageResponse{Message: s}, u.Connection)
// }
