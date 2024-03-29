package router

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/game/admin"
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
}

func (r *Router) Disconnect(c *server.Connection) {
	delete(connections, c.ID)
	r.game.Disconnect <- c.ID
}

func (r *Router) Handle(m *server.InboundMessage) {
	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.Payload, &obj)
	if err != nil {
		// invalid JSON format?
		logger.Write.Error("Invalid JSON formatting")
		return
	}

	command, commandExists := obj["command"]

	if !commandExists {
		// no comand found in JSON payload, invalid JSON then
		logger.Write.Error("No command found in JSON payload")
		return
	}

	args, argsExist := obj["args"]

	var cmd string
	err = json.Unmarshal(*command, &cmd)

	c := connections[m.Connection.ID]

	if !argsExist {
		if c.Character != nil {
			logger.Write.Warning("No args found in JSON payload for command: %s for connection %s; continuing", cmd, m.Connection.ID)
		}

		args = &json.RawMessage{}
	}

	lowered := strings.ToLower(cmd)
	method := strings.ToUpper(lowered[:1]) + lowered[1:]

	if len(cmd) > 6 && c.Socket.User.IsAdmin() && cmd[:6] == "admin_" {
		adminCommand := reflect.ValueOf(&admin.Controller).MethodByName(strings.ToUpper(cmd[6:7]) + cmd[7:])
		if adminCommand.IsValid() {
			adminCommand.Call([]reflect.Value{reflect.ValueOf(c), reflect.ValueOf(r.game), reflect.ValueOf(args)})
			return // stop execution here
		}
	}

	if c.Character != nil {
		r.game.Input <- &game.Command{Value: lowered, Args: args, Character: c.Character, Connection: c.Socket}
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
