package engine

import "gopkg.in/mgo.v2/bson"

func Bootstrap() {
	BootstrapSuperAdmin()
}

func BootstrapSuperAdmin() {
	errors := CreateUser("superadmin@minimiri.com", "superadmin")
	if len(errors) > 0 {
		logger.Info("Admin already existed.")
		return
	}

	q := bson.M{"email": "superadmin@minimiri.com"}
	change := bson.M{"$set": bson.M{"isadmin": true}}
	err := db.C("users").Update(q, change)
	if err != nil {
		panic(err)
	}

	logger.Info("Created Super Admin user.\n")
}
