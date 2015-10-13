package engine

import (
	db "github.com/jonathonharrell/miri-ws-server/engine/core/database"
	"github.com/jonathonharrell/miri-ws-server/engine/logger"
	"gopkg.in/mgo.v2/bson"
)

func Bootstrap() {
	BootstrapSuperAdmin()
}

func BootstrapSuperAdmin() {
	errors := CreateUser("superadmin@minimiri.com", "superadmin", &User{})
	if len(errors) > 0 {
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
