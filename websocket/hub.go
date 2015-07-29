package websocket

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	Connections map[*Connection]bool

	// Inbound messages from the connections.
	Broadcast chan []byte

	// Register requests from the connections.
	Register chan *Connection

	// Unregister requests from connections.
	Unregister chan *Connection

	OnConnect func(c *Connection)
	OnMessage func(c *Connection)
	OnDisconnect func(c *Connection)
}

var h = Hub{
	Broadcast:   make(chan []byte),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	Connections: make(map[*Connection]bool),
}

func (h *Hub) run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c] = true
			h.OnConnect(c)
		case c := <-h.Unregister:
			if _, ok := h.Connections[c]; ok {
				delete(h.Connections, c)
				close(c.Send)
			}
		case m := <-h.Broadcast:
			for c := range h.Connections {
				select {
				case c.Send <- m:
				default:
					close(c.Send)
					delete(h.Connections, c)
				}
			}
		}
	}
}

func (h *Hub) SetOnConnectCallback(callback func(c *Connection)) {
	h.OnConnect = callback
}

func (h *Hub) SetOnMessageCallback(callback func(c *Connection)) {
	h.OnMessage = callback
}

func (h *Hub) SetOnDisconnectCallback(callback func(c *Connection)) {
	h.OnDisconnect = callback
}
