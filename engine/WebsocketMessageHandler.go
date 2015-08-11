package engine

import (
	"encoding/json"
	"fmt"
)

type errorHandler func(u *User, args ...interface{})

var (
	MethodSet           map[int]map[string]func(u *User, args *json.RawMessage)
	InvalidStateHandler errorHandler
	InvalidHandlerIndex errorHandler
)

func InitMessageHandler() {
	MethodSet = make(map[int]map[string]func(u *User, args *json.RawMessage))

	MethodSet[NotAuthenticated] = make(map[string]func(u *User, args *json.RawMessage))
	MethodSet[Authenticated] = make(map[string]func(u *User, args *json.RawMessage))
	MethodSet[InGame] = make(map[string]func(u *User, args *json.RawMessage))
}

// @todo, this should return an error if the key is already defined
func AddHandler(state int, name string, handler func(u *User, args *json.RawMessage)) {
	MethodSet[state][name] = handler
}

func Route(name string, u *User, args *json.RawMessage) {
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

func Interpret(m *Message, u *User) {
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
