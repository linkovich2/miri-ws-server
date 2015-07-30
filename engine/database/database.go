package database

import (
  "fmt"
  "gopkg.in/mgo.v2"
)

type DB struct {
  session *mgo.Session
}

func Connect() *DB {
  s, err := mgo.Dial("localhost:27017") // @temp this should connect dynamically based on environment variables with sane defaults
  if err != nil {
    panic(err)
  }

  defer s.Close()

  s.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.
  
  return &DB{session: s}
}
