package core

type (
	Realm struct {
		ID    string
		Name  string `json:"name"`
		Rooms map[string]*Room
		// Weathers  map[string]Weather
		// TimeCycle map[string]RealmTime `json:"times"`
		// Time      string
	}

	RealmTime struct {
		Name        string
		Min         int    // minute during hour this time will display
		Description string // the lighting, smells, etc.
	}
)

func (r *Realm) Update(sendMsg func(string, string)) {
	for _, r := range r.Rooms {
		r.Update(sendMsg)
	}
}
