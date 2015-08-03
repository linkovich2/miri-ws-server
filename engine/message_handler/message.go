// Package message contains...
package message_handler

import (
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"fmt"
)

type errorHandler func(u *auth.User, args ...interface{})

var (
	MethodSet map[int]map[string]func(u *auth.User, args ...interface{})
	InvalidStateHandler errorHandler
	InvalidHandlerIndex errorHandler
)

func Init() {
	MethodSet = make(map[int]map[string]func(u *auth.User, args ...interface{}))

	MethodSet[auth.NotAuthenticated] = make(map[string]func(u *auth.User, args ...interface{}))
	MethodSet[auth.Authenticated]    = make(map[string]func(u *auth.User, args ...interface{}))
	MethodSet[auth.InGame]           = make(map[string]func(u *auth.User, args ...interface{}))
}

// @todo, this should return an error if the key is already defined
func AddHandler(state int, name string, handler func(u *auth.User, args ...interface{})) {
	MethodSet[state][name] = handler
}

func Route(name string, u *auth.User, args ...interface{}) {
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
	fmt.Printf("%v\n", u.Connection.ID)

	var n int
	for i := 0; i < len(m.Payload); i++ {
		if string(m.Payload[i]) == " " {
			n = i
			break
		}
	}

	cmd := string(m.Payload[0:n])
	args := string(m.Payload[(n + 1):len(m.Payload)])
	Route(cmd, u, args)
}
