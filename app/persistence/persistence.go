package persistence

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
)

func SaveCharacter(c *core.Character) {
	go func() {
		session, dbName := database.GetSession() // connect
		db := session.DB(dbName)
		defer session.Close()

		copy := *c
		copy.RemoveState(core.StateLoggingOut)
		copy.RemoveState(core.StateMoving)

		err := db.C("characters").UpdateId(c.ID, &copy)
		if err != nil {
			logger.Write.Error(err.Error())
		}
	}()
}
