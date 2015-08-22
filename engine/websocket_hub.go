package engine

import "encoding/json"

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type (
	ConnectionHub struct {
		Connections map[*Connection]bool // Registered connections.
		inbound     chan *Message        // Inbound messages from the connections.
		register    chan *Connection     // Register requests from the connections.
		unregister  chan *Connection     // Unregister requests from connections.
	}

	MessageResponse struct {
		ResponseTo string      `json:"response_to"`
		Errors     []string    `json:"errors"`
		Success    bool        `json:"success"`
		Message    string      `json:"message"` // @todo this should probably be another struct with specific areas to deliver messages to in the UI
		Data       interface{} `json:"data"`
	}
)

var hub = ConnectionHub{
	Connections: make(map[*Connection]bool),
	inbound:     make(chan *Message),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
}

func (h *ConnectionHub) run() {
	for {
		select {
		case c := <-h.register:
			h.Connections[c] = true

			logger.Info("New Connection [%s]", c.ID)
			users[c.ID] = &User{Connection: c, State: NotAuthenticated}
			// @todo FUTURE FEATURE we should also try to authenticate here based on cookie / json web token / session

		case c := <-h.unregister:
			if _, ok := h.Connections[c]; ok {
				// @todo FUTURE FEATURE we need to remove the session from the DB,
				// for now we just close out the connection and it is no longer authenticated

				// run any other logic on disconnect we need here
				logger.Notice("Connection [%s] disconnected", c.ID)

				delete(h.Connections, c)
				close(c.send)
			}
		case m := <-h.inbound:
			interpret(m, users[m.Connection.ID])
		}
	}
}

// Send a message to a lot of connections
func (h *ConnectionHub) Broadcast(m *MessageResponse, targets []*Connection) {
	msg, _ := json.Marshal(m)
	for _, c := range targets {
		c.send <- msg
	}
}

// Send a message to one connection
func (h *ConnectionHub) Send(m *MessageResponse, c *Connection) {
	msg, _ := json.Marshal(m)
	c.send <- msg
}
