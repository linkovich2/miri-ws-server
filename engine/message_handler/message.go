// Package message contains...
package message_handler

import (
	"bytes"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
	"reflect"
	"strings"
)

func Route(name string, args ...interface{}) {
	var m *websocket.Message
	if len(args) > 0 {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		reflect.ValueOf(m).MethodByName(name).Call(inputs)
	} else if len(args) == 0 {
		reflect.ValueOf(m).MethodByName(name).Call([]reflect.Value{})
	} else if reflect.ValueOf(m).NumField() == 0 {
		// reflect.ValueOf(m).MethodByName("Say").Call(reflect.Value{"That isn't a command!"})
	}
}

func Interpreter(a *websocket.Message) {
	n := bytes.Index(a.Payload, []byte{0})
	str := string(a.Payload[:n])
	firstWord := strings.Fields(str)
	Route(firstWord[0], str[1:n])
}
