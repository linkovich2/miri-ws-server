package game

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/server"
)

type Command struct {
	Value      string
	Args       *json.RawMessage
	Character  *core.Character
	Connection *server.Connection
}

var commandRegistrar = map[string]func(*Game, *Command){
	"move": cMove,
}
