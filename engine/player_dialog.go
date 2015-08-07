package engine

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

type ChatArgs struct {
	Input string
}

func CmdSay(u *auth.User, args *json.RawMessage) {
	chat := &ChatArgs{}
	json.Unmarshal(*args, &chat)

	s := filters.ReplaceProfanity(chat.Input)
	// @todo RP filter and any other filters we want chat to go through

	hub.Send([]byte(s), u.Connection)
}
