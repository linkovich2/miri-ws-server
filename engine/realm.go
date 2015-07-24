package engine

import (

)

const (
  MAX_ROOMS_PER_REALM = 512 // max number of rooms per realm
  MAX_WEATHER         = 16 // max number of "storms" that can be "on" at any given time (per realm)
)

type (
  Realm struct {
    ID        int
    Name      string               `json:"name"`
    Rooms     map[string]Room
    Weathers  map[string]Weather
    TimeCycle map[string]RealmTime `json:"times"`
    Time      string
  }

  RealmTime struct {
    Name        string
    Min         int     // minute during hour this time will display
    Description string  // the lighting, smells, etc.
  }
)

func (realm *Realm) Init() {
  // @todo call FromJSON here, build the length of TimeCycle from JSON arr length
  // build the rooms

  realm.Rooms    = make(map[string]Room, MAX_ROOMS_PER_REALM) // init Room map
  realm.Weathers = make(map[string]Weather, MAX_WEATHER)     // init Weathers
}


func (realm *Realm) Update() {
  for _, r := range realm.Rooms {
    r.Update()
  }

  // also:
  // generate weather
  // update realm time
}


func (realm *Realm) FromJSON(filepath string) {
  // create a realm from a json file
}
