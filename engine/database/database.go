package database

import (
  "fmt"
  "gopkg.in/mgo.v2"
)

type DB struct {
  session *mgo.Session
  Collections map[string]*mgo.Collection
}

func (db *DB) Init() {

}

func (db *DB) Connect(host string, database string) {
  s, err := mgo.Dial(host)
  if err != nil {
    panic(err)
  }

  defer s.Close()

  s.SetMode(mgo.Monotonic, true) // Optional. Switch the session to a monotonic behavior.

  db.session = s.DB(database)
}

func (db *DB) RegisterCollection(name string) {
  db.Collections[name] = db.session.C(name)
}
