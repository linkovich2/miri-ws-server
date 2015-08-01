package database

import (
  "gopkg.in/mgo.v2"
)

var (
  session *mgo.Session
  DB *mgo.Database
)

func Connect(host string, database string) {
  s, err := mgo.Dial(host)
  if err != nil {
    panic(err)
  }

  session = s

  session.SetMode(mgo.Monotonic, true) // @resource http://godoc.org/labix.org/v2/mgo#Session.SetMode

  DB = session.DB(database)
}

func Close() {
  session.Close()
}
