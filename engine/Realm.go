package engine

const (
	maxRoomsPerRealm = 512 // max number of rooms per realm
	maxWeather       = 16  // max number of "storms" that can be "on" at any given time (per realm)
)

type (
	realm struct {
		id        int
		name      string `json:"name"`
		rooms     map[string]room
		weathers  map[string]weather
		timeCycle map[string]realmTime `json:"times"`
		time      string
	}

	realmTime struct {
		name        string
		min         int    // minute during hour this time will display
		description string // the lighting, smells, etc.
	}
)

func (r *realm) init() {
	// @todo call FromJSON here, build the length of TimeCycle from JSON arr length
	// build the rooms

	r.rooms = make(map[string]room, maxRoomsPerRealm) // init Room map
	r.weathers = make(map[string]weather, maxWeather) // init Weathers
}

func (r *realm) update() {
	for _, r := range r.rooms {
		r.update()
	}

	// also:
	// generate weather
	// update realm time
}
