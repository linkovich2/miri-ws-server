package engine

import (
	"encoding/json"
	"reflect"
	"strings"
)

type (
	errorHandler     func(u *user, args ...interface{})
	handlerInterface struct{}
)

var (
	handlers = &handlerInterface{}
	aliases  = make(map[string]string)

	invalidStateHandler errorHandler
	invalidHandlerIndex errorHandler
)

func addAlias(alt string, cmd string) {
	aliases[alt] = cmd
}

func route(name string, u *user, args *json.RawMessage) {
	var method string

	// capitalize first letter of command
	c := strings.ToUpper(name)

	if alias, exists := aliases[c]; exists {
		method = strings.Join([]string{"Command", stateString(u.state), "_", alias}, "")
	} else {
		method = strings.Join([]string{"Command", stateString(u.state), "_", c}, "")
	}

	cmd := reflect.ValueOf(handlers).MethodByName(method)

	if cmd.IsValid() {
		cmd.Call([]reflect.Value{reflect.ValueOf(u), reflect.ValueOf(args)})
	} else {
		invalidHandlerIndex(u, args)
	}
}

func interpret(m *message, u *user) {
	// @todo we should log the received command
	// @todo we should probably be able to attach a logger to the message handler

	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.payload, &obj)
	if err != nil {
		// invalid JSON format?
		logger.Error("Invalid JSON formatting") // @todo errors
		return
	}

	command, commandExists := obj["command"]

	if !commandExists {
		// no comand found in JSON payload, invalid JSON then
		logger.Error("No command found in JSON payload") // @todo error handling
		return
	}

	args, argsExist := obj["args"]

	var cmd string
	err = json.Unmarshal(*command, &cmd)

	if !argsExist {
		logger.Warning("No args found in JSON payload for command: %s for connection %s; continuing", cmd, u.connection.id)
		args = &json.RawMessage{}
	}

	route(cmd, u, args)
}
