// Package message contains...
package message_handler

import (
	"fmt"
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

type errorHandler func(u *auth.User, args ...interface{})

var (
	MethodSet           map[int]map[string]func(u *auth.User, args ...interface{})
	InvalidStateHandler errorHandler
	InvalidHandlerIndex errorHandler
)

func Init() {
	MethodSet = make(map[int]map[string]func(u *auth.User, args ...interface{}))

	MethodSet[auth.NotAuthenticated] = make(map[string]func(u *auth.User, args ...interface{}))
	MethodSet[auth.Authenticated] = make(map[string]func(u *auth.User, args ...interface{}))
	MethodSet[auth.InGame] = make(map[string]func(u *auth.User, args ...interface{}))
}

// @todo, this should return an error if the key is already defined
func AddHandler(state int, name string, handler func(u *auth.User, args ...interface{})) {
	MethodSet[state][name] = handler
}

func Route(name string, u *auth.User, args *json.RawMessage) {
	if s, ok := MethodSet[u.State]; ok {
		if cmd, exists := s[name]; exists {
			cmd(u, args) // we're all good, yay!
		} else {
			InvalidHandlerIndex(u, args)
		}
	} else {
		InvalidStateHandler(u, args)
	}
}

func Interpret(m *websocket.Message, u *auth.User) {
	// @todo we should log the received command
	// @todo we should probably be able to attach a logger to the message handler

	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.Payload, &obj)
	if err != nil {
		// invalid JSON format?
		fmt.Println("Invalid JSON formatting") // @todo errors
		return
	}

	command, commandExists := obj["command"]

	if !commandExists {
		// no comand found in JSON payload, invalid JSON then
		fmt.Println("No command found in JSON payload") // @todo error handling
		return
	}

	args, argsExist := obj["args"]

	if !argsExist {
		fmt.Println("No args found in JSON payload; continuing") // @todo error handling
		args = &json.RawMessage{}
	}

	var cmd string
	err = json.Unmarshal(*command, &cmd)

	Route(cmd, u, args)
}
