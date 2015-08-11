package engine

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

type chatArgs struct {
	Input string
}

func (h *handlerInterface) CmdSay(u *user, args *json.RawMessage) {
	logger.Info("Called 'say'!")
	chat := &chatArgs{}
	json.Unmarshal(*args, &chat)

	s := filters.ReplaceProfanity(chat.Input)
	// @todo RP filter and any other filters we want chat to go through

	hub.Send([]byte(s), u.connection)
}
