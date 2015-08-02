// Package message contains...
package message_handler

import (
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
	"reflect"
	"fmt"
)

type Methods struct {}
var m = &Methods{}

func (m *Methods) Say(args ...interface{}) {
	fmt.Printf("%v\n", args)
}

func Route(name string, args ...interface{}) {
	if len(args) > 0 {
		inputs := make([]reflect.Value, len(args))
		for i, _ := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		reflect.ValueOf(m).MethodByName(name).Call(inputs)
	}// else if len(args) == 0 {
	// 	reflect.ValueOf(m).MethodByName(name).Call([]reflect.Value{})
	// } else if reflect.ValueOf(m).NumField() == 0 {
	// 	// reflect.ValueOf(m).MethodByName("Say").Call(reflect.Value{"That isn't a command!"})
	// }
}

func Interpreter(m *websocket.Message) {
	var n int
	for i := 0; i < len(m.Payload); i++ {
		if string(m.Payload[i]) == " " {
			n = i
			break
		}
	}

	cmd := string(m.Payload[0:n])
	args := string(m.Payload[(n + 1):len(m.Payload)])
	Route(cmd, args)
}
