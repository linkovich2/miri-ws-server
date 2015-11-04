package websockets

import (
	"net/http"
	"stablelib.com/v1/uniuri"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
)

func WebsocketServe(w http.ResponseWriter, r *http.Request) {
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    logger.Write.Error(err.Error())
    return
  }

  c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New()}
  Hub.register <- c
  go c.writePump()
  c.readPump()
}
