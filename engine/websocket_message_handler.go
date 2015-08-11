package engine

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type (
	errorHandler func(u *user, args ...interface{})
	handlerInterface struct {}

	aliasList map[string]string
)

var (
	handlers = &handlerInterface{}
	aliases = make(map[string]string)

	invalidStateHandler errorHandler
	invalidHandlerIndex errorHandler
)

func (a *aliasList) AddAlias(alt string, cmd string) {
	a[alt] = cmd
}

func route(name string, u *user, args *json.RawMessage) {
	var method string

	// because we'd need to call public methods on the handler only
	// and it's easier and faster to concatenate strings then to capitalize the first letter
	if alias, exists := aliases[name]; exists {
		method = strings.Join([]string{"Cmd_", alias}, "")
	} else {
		method := strings.Join([]string{"Cmd_", name}, "")
	}

	cmd := reflect.ValueOf(handlers).MethodByName(method)

	// @todo check if method exists first
	// if cmd.isValid() {
		cmd.Call([]reflect.Value{reflect.ValueOf(u), reflect.ValueOf(args)})
	// } else {
	// 	invalidHandlerIndex(u, args)
	// }
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
