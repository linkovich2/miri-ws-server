package engine

func attachMessageHandlers() {
	InitMessageHandler() // set up message handler and router

	// -- BEGIN register message handlers
	AddHandler(NotAuthenticated, "say", cmdSay)
	AddHandler(NotAuthenticated, "authenticate", cmdAuthenticate)
	// -- END

	// -- BEGIN register error handlers
	InvalidStateHandler = func(u *User, args ...interface{}) {
		hub.Send([]byte("State not valid for some reason."), u.Connection)
	}

	InvalidHandlerIndex = func(u *User, args ...interface{}) {
		hub.Send([]byte("Command not found."), u.Connection)
	}
	// -- END
}
