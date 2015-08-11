package engine

func attachMessageHandlers() {
	invalidStateHandler = func(u *user, args ...interface{}) {
		hub.Send([]byte("State not valid for some reason."), u.connection)
	}

	invalidHandlerIndex = func(u *user, args ...interface{}) {
		hub.Send([]byte("Command not found."), u.connection)
	}
}
