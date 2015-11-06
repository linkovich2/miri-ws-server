package server

type (
	ConnectionHandler interface {
		Connect(c *Connection)
		Disconnect(c *Connection)
		Handle(m *InboundMessage)
	}
)

type defaultConnectionHandler struct{}

func (d *defaultConnectionHandler) Connect(c *Connection) {
	return
}

func (d *defaultConnectionHandler) Disconnect(c *Connection) {
	return
}

func (d *defaultConnectionHandler) Handle(m *InboundMessage) {
	return
}
