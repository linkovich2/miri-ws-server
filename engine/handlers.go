package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/auth"
	"github.com/jonathonharrell/miri-ws-server/engine/message_handler"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

func attachMessageHandlers() {
	// pass global hub to subpackages
	makeHubAvailable()

	hub.SetOnConnectCallback(func(c *websocket.Connection) {
		users[c.ID] = &auth.User{Connection: c, State: auth.NotAuthenticated}
		// maybe we should also try to authenticate, if we want to use cookies or whatever
	})

	hub.SetOnMessageCallback(func(m *websocket.Message) {
		message_handler.Interpret(m, users[m.Connection.ID])
	})

	message_handler.Init() // set up message handler and router

	// -- BEGIN register message handlers
	message_handler.AddHandler(auth.NotAuthenticated, "say", CmdSay)
	message_handler.AddHandler(auth.NotAuthenticated, "authenticate", auth.CmdAuthenticate)
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

func makeHubAvailable() {
	auth.SetHub(hub)
}
