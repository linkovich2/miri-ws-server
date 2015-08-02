// Package message contains...
package message

import (
	"bytes"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
	"reflect"
	"strings"
)

func Route(any interface{}, name string, args ...interface{}) {
	var m *websocket.Message
	if len(args) > 0 {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		reflect.ValueOf(any).MethodByName(name).Call(inputs)
	} else if args == 0 {
		reflect.ValueOf(any).MethodByName(name).Call([]reflect.Value{})
	} else if reflect.ValueOf(any).NumField() == 0 {
		reflect.ValueOf(m).MethodByName("Say").Call([]reflectValue{"That isn't a command!"})
	}
	return
}

func Interpreter(any interface{}, a []byte) {
	n := bytes.Index(a, []byte{0})
	str := string(a[:n])
	firstWord := strings.Fields(str)
	Route(any, firstWord[0], str[1:n])
	return
}
