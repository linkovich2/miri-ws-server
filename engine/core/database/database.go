package database

import (
	"gopkg.in/mgo.v2"
)

var (
	db      *mgo.Database
	session *mgo.Session
)

func ConnectToDatabase(host string, database string) {
	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session = s
	session.SetMode(mgo.Monotonic, true) // @resource http://godoc.org/labix.org/v2/mgo#Session.SetMode

	db = session.DB(database)
}

func CloseDatabaseConnection() {
	session.Close()
}

func GetDB() *mgo.Database {
	return db
}
