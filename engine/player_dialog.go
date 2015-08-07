package engine

import (
	"encoding/json"

	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/util/filters"
)

type ChatArgs struct {
	Input string
}

// @todo this should go through a profanity and RP filter
// Replace List:
// all curse words: IE, fuck shit cunt etc
// as many acronyms as possible: lol -> Haha!, rofl -> HAHAHA!
func CmdSay(u *auth.User, args *json.RawMessage) {
	chat := &ChatArgs{}
	json.Unmarshal(*args, &chat)

	s := filters.ProfanityFilter(chat.Input)

	hub.Send([]byte(s), u.Connection)
}
