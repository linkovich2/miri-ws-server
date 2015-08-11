package engine

import (
	"encoding/json"
	"fmt"
)

type errorHandler func(u *user, args ...interface{})

var (
	methodSet           map[int]map[string]func(u *user, args *json.RawMessage)
	invalidStateHandler errorHandler
	invalidHandlerIndex errorHandler
)

func initMessageHandler() {
	methodSet = make(map[int]map[string]func(u *user, args *json.RawMessage))

	methodSet[notAuthenticated] = make(map[string]func(u *user, args *json.RawMessage))
	methodSet[authenticated] = make(map[string]func(u *user, args *json.RawMessage))
	methodSet[inGame] = make(map[string]func(u *user, args *json.RawMessage))
}

// @todo, this should return an error if the key is already defined
func addHandler(state int, name string, handler func(u *user, args *json.RawMessage)) {
	methodSet[state][name] = handler
}

func route(name string, u *user, args *json.RawMessage) {
	if s, ok := methodSet[u.state]; ok {
		if cmd, exists := s[name]; exists {
			cmd(u, args) // we're all good, yay!
		} else {
			invalidHandlerIndex(u, args)
		}
	} else {
		invalidStateHandler(u, args)
	}
}

func interpret(m *message, u *user) {
	// @todo we should log the received command
	// @todo we should probably be able to attach a logger to the message handler

	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.payload, &obj)
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

	route(cmd, u, args)
}
