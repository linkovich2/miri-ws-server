package database

import (
	"gopkg.in/mgo.v2"
)

var (
	name    string
	session *mgo.Session
)

func ConnectToDatabase(host, db string) {
	name = db

	s, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}

	session = s
	session.SetMode(mgo.Monotonic, true) // @resource http://godoc.org/labix.org/v2/mgo#Session.SetMode
}

func CloseDatabaseConnection() {
	session.Close()
}

func GetSession() (*mgo.Session, string) {
	return session.Copy(), name
}
