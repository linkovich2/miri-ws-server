package engine

import (
	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/services"
	"github.com/jonathonharrell/miri-ws-server/engine/api/parameters"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"gopkg.in/mgo.v2/bson"
)

func Bootstrap() {
	BootstrapSuperAdmin()
}

func BootstrapSuperAdmin() {
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
