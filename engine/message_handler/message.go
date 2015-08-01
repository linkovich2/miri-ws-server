// Package websocket contains...
package websocket

import (
	"bytes"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket/connection.go"
	"reflect"
	"strconv"
	"strings"
)

func (m *Message) Route(name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(m).MethodByName(name).Call(inputs)
	return
}

func (m *Message) Interpreter(a []byte) {
	n := bytes.Index(a, []byte{0})
	str := string(a[:n])
	strings.Split(str, " ")
	m.Route(strconv.Itoa(int(str[0])), str[1])
	return
}
