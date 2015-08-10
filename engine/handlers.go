package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/message_handler"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

func attachMessageHandlers() {
	hub.SetOnConnectCallback(onConnectCallback)

	hub.SetOnMessageCallback(func(m *websocket.Message) {
		message_handler.Interpret(m, users[m.Connection.ID])
	})

	message_handler.Init() // set up message handler and router

	// -- BEGIN register message handlers
	message_handler.AddHandler(auth.NotAuthenticated, "say", cmdSay)
	message_handler.AddHandler(auth.NotAuthenticated, "authenticate", cmdAuthenticate)
	// -- END

	// -- BEGIN register error handlers
	message_handler.InvalidStateHandler = func(u *auth.User, args ...interface{}) {
		hub.Send([]byte("State not valid for some reason."), u.Connection)
	}

	message_handler.InvalidHandlerIndex = func(u *auth.User, args ...interface{}) {
		hub.Send([]byte("Command not found."), u.Connection)
	}
	// -- END
}
