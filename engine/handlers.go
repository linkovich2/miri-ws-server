package engine

func attachMessageHandlers() {
	initMessageHandler() // set up message handler and router

	// -- BEGIN register message handlers
	addHandler(notAuthenticated, "say", cmdSay)
	addHandler(notAuthenticated, "authenticate", cmdAuthenticate)
	// -- END

	// -- BEGIN register error handlers
	invalidStateHandler = func(u *user, args ...interface{}) {
		hub.Send([]byte("State not valid for some reason."), u.connection)
	}

	invalidHandlerIndex = func(u *user, args ...interface{}) {
		hub.Send([]byte("Command not found."), u.connection)
	}
	// -- END
}
