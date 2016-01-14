package game

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/server"
)

type Command struct {
	Value      string
	Args       *json.RawMessage
	Character  *core.Character
	Connection *server.Connection
}

type cBasicArgs struct {
	Input string `json:"input"`
}

var commandRegistrar = map[string]func(*Game, *Command){
	"move":     cMove,
	"go":       cMove,
	"walk":     cMove,
	"run":      cMove,
	"say":      cSay,
	"yell":     cYell,
	"target":   cAddTarget,
	"untarget": cRemoveTarget,
	"interact": cInteract,
}

func (c *Command) GetInput() (string, error) {
	params := cBasicArgs{}
	err := json.Unmarshal(*c.Args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return "", errors.New(fmt.Sprintf("Could not gather input from string [%s]", string(*c.Args)))
	}

	return params.Input, nil
}
