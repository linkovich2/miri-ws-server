package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/server"
)

type Connection struct {
	Socket    *server.Connection
	Character *core.Character
}
