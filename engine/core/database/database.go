package database

import (
	"gopkg.in/mgo.v2"
	"github.com/jonathonharrell/miri-ws-server/engine/settings"
)

var (
  env settings.Environment
	session *mgo.Session
)

func ConnectToDatabase() {
	env = settings.GetEnv()

	s, err := mgo.Dial(env.DBHost)
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
	return session.Copy(), env.DBName
}
