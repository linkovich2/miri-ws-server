package database

import (
  "gopkg.in/mgo.v2"
)

func Connect(host string, database string) *mgo.Database {
  s, err := mgo.Dial(host)
  if err != nil {
    panic(err)
  }

  defer s.Close()

  s.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.

  return s.DB(database)
}
