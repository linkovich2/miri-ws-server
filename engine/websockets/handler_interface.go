package websockets

type (
	ConnectionHandler interface {
		Connect(c *Connection) bool
		Disconnect(c *Connection) bool
		Handle(m *Message) bool
	}
)

type DefaultConnectionHandler struct{}

func (d *DefaultConnectionHandler) Connect(c *Connection) bool {
	return false
}

func (d *DefaultConnectionHandler) Disconnect(c *Connection) bool {
	return false
}

func (d *DefaultConnectionHandler) Handle(m *Message) bool {
	return false
}
