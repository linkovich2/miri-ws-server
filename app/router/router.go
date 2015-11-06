package router

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/game/characters"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/server"
	"reflect"
	"strings"
)

type (
	Router struct {
		aliases map[string]string
		game    *game.Game
	}
)

var connections = make(map[string]*game.Connection)

func (r *Router) Connect(c *server.Connection) {
	connections[c.ID] = &game.Connection{Socket: c}
	// do we even need to do anything on connect? @todo
}

func (r *Router) Disconnect(c *server.Connection) {
	delete(connections, c.ID)

	// @todo we should also pass this to the game goroutine,
	// and it should check if the connection is in-game
	// if it is, it should remove the connection after a set time period
}

func (r *Router) Handle(m *server.InboundMessage) {
	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.Payload, &obj)
	if err != nil {
		// invalid JSON format?
		logger.Write.Error("Invalid JSON formatting") // @todo errors
		return
	}

	command, commandExists := obj["command"]

	if !commandExists {
		// no comand found in JSON payload, invalid JSON then
		logger.Write.Error("No command found in JSON payload") // @todo error handling
		return
	}

	args, argsExist := obj["args"]

	var cmd string
	err = json.Unmarshal(*command, &cmd)

	if !argsExist {
		logger.Write.Warning("No args found in JSON payload for command: %s for connection %s; continuing", cmd, m.Connection.ID)
		args = &json.RawMessage{}
	}

	c := connections[m.Connection.ID]
	method := strings.ToLower(cmd)
	method = strings.ToUpper(method[:1]) + method[1:]

	if c.Character != nil {
		r.game.Input <- &game.Command{Value: method, Args: args, Character: c.Character, Connection: c.Socket}
	} else {
		command := reflect.ValueOf(&characters.Controller).MethodByName(method)

		if command.IsValid() {
			command.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(r.game), reflect.ValueOf(args)})
		} else {
			logger.Write.Error("Handler not found for %v", method)
		}
	}
}

// Add an alias to the list
func (r *Router) AddAlias(alt string, cmd string) {
	r.aliases[alt] = cmd
}

func NewRouter(g *game.Game) *Router {
	return &Router{make(map[string]string), g}
}
