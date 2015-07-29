package websocket

import (
	"log"
)

/*hub passes the messages through this interface and we **Do-Something** with them*/

type msgHandler struct {
	recieve chan []byte
	send    chan []byte
}

var msgH = &msgHandler{
	recieve: make(chan []byte),
	send:    make(chan []byte),
}

func (m *msgHandler) OnMessage(str string) {
	log.Println(str)
}
