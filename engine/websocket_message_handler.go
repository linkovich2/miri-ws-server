package engine

import (
	"encoding/json"
	"reflect"
	"strings"
)

type HandlerInterface struct{} // empty struct to attach message handlers to

var (
	handlers = &HandlerInterface{}
	aliases  = make(map[string]string)
)

func handlerNotFoundError(u *User, args ...interface{}) {
	hub.Send(&MessageResponse{Errors: []string{"Either wrong state or command does not exist."}}, u.Connection)
}

// Add an alias to the list
func AddAlias(alt string, cmd string) {
	aliases[alt] = cmd
}

func routeToCommand(name string, u *User, args *json.RawMessage) {
	logger.Info("Connection [%s]: Processing %s command...", u.Connection.ID, name)

	var method string

	// capitalize first letter of command
	c := strings.ToUpper(name)

	if alias, exists := aliases[c]; exists {
		method = strings.Join([]string{"Command", StateString(u.State), "_", alias}, "")
	} else {
		method = strings.Join([]string{"Command", StateString(u.State), "_", c}, "")
	}

	cmd := reflect.ValueOf(handlers).MethodByName(method)

	if cmd.IsValid() {
		cmd.Call([]reflect.Value{reflect.ValueOf(u), reflect.ValueOf(args)})
	} else {
		logger.Error(" -- '%s' command wasn't found for the given state!", name)
		handlerNotFoundError(u, args)
	}
}

func interpret(m *Message, u *User) {
	// @todo we should log the received command
	// @todo we should probably be able to attach a logger to the message handler

	var obj map[string]*json.RawMessage
	err := json.Unmarshal(m.Payload, &obj)
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
		args = &json.RawMessage{}
	}

	routeToCommand(cmd, u, args)
}
