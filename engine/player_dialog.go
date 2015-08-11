package engine

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

type chatArgs struct {
	input string
}

func cmdSay(u *user, args *json.RawMessage) {
	chat := &chatArgs{}
	json.Unmarshal(*args, &chat)

	s := filters.ReplaceProfanity(chat.input)
	// @todo RP filter and any other filters we want chat to go through

	hub.Send([]byte(s), u.connection)
}
