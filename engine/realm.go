package engine

import ()

const (
	MaxRoomsPerRealm = 512 // max number of rooms per realm
	MaxWeather       = 16  // max number of "storms" that can be "on" at any given time (per realm)
)

type (
	Realm struct {
		ID        int
		Name      string `json:"name"`
		Rooms     map[string]Room
		Weathers  map[string]Weather
		TimeCycle map[string]RealmTime `json:"times"`
		Time      string
	}

	RealmTime struct {
		Name        string
		Min         int    // minute during hour this time will display
		Description string // the lighting, smells, etc.
	}
)

func (r *Realm) Init() {
	// @todo call FromJSON here, build the length of TimeCycle from JSON arr length
	// build the rooms

	r.Rooms = make(map[string]Room, MaxRoomsPerRealm) // init Room map
	r.Weathers = make(map[string]Weather, MaxWeather) // init Weathers
}

func (r *Realm) Update() {
	for _, r := range r.Rooms {
		r.Update()
	}

	// also:
	// generate weather
	// update realm time
}
