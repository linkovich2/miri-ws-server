package engine

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type connectionHub struct {
	connections map[*connection]bool // Registered connections.
	inbound     chan *message        // Inbound messages from the connections.
	register    chan *connection     // Register requests from the connections.
	unregister  chan *connection     // Unregister requests from connections.
}

var hub = connectionHub{
	connections: make(map[*connection]bool),
	inbound:     make(chan *message),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
}

func (h *connectionHub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			users[c.id] = &user{connection: c, state: notAuthenticated}
			// maybe we should also try to authenticate, if we want to use cookies or whatever
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				// run any other logic on disconnect we need here
				logger.Notice("Connection [%s] disconnected", c.id)

				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			interpret(m, users[m.connection.id])
		}
	}
}

// Send a message to a lot of connections
func (h *connectionHub) Broadcast(msg []byte, targets []*connection) {
	for _, c := range targets {
		c.send <- msg
	}
}

// Send a message to one connection
func (h *connectionHub) Send(msg []byte, c *connection) {
	c.send <- msg
}
