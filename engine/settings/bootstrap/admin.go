package bootstrap

import (
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"github.com/jonathonharrell/miri-ws-server/engine/services"
	"gopkg.in/mgo.v2/bson"
)

func bootstrapSuperAdmin() {
	status, _ := services.CreateUser(&parameters.User{"superadmin@minimiri.com", "superadmin"})
	if status != 200 { // we got an error back, likely that the admin account already exists
		logger.Write.Info("Admin already existed.")
		return
	}

	q := bson.M{"email": "superadmin@minimiri.com"}
	change := bson.M{"$set": bson.M{"isadmin": true}}
	err := db.GetDB().C("users").Update(q, change)
	if err != nil {
		panic(err)
	}

	logger.Write.Info("Created Super Admin user.\n")
}
