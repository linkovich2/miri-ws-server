package websockets

import (
	"github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"stablelib.com/v1/uniuri"
)

type ConnectionData struct {
	User      *models.User
	Character *models.Character
}

func WebsocketServe(w http.ResponseWriter, r *http.Request, userId string) {
	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	u := models.User{}
	err := db.C("users").Find(bson.M{"_id": bson.ObjectIdHex(userId)}).One(&u)

	if err != nil { // no existing user
		logger.Write.Error(err.Error())
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	// @todo attach character object to connection
	c := &Connection{send: make(chan []byte, 256), webSocket: ws, ID: uniuri.New(), IsAdmin: u.IsAdmin}
	Hub.register <- c
	go c.writePump()
	c.readPump()
}
