package core

import "encoding/json"

var races map[string]Race

type Race struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
	ID              string `json:"id, omitempty"`
	Descriptor      string `json:"descriptor"`
	GenderHuman     bool   `json:"gender_human"`
}

// Content methods
func GetRaces() map[string]Race {
	if len(races) <= 0 {
		data, err := Asset("json/races.json")
		if err != nil {
			panic(err)
		}

		a := map[string]Race{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		races = a
	}

	return races
}

func GetRace(key string) Race {
	races := GetRaces()
	return races[key]
}

// FUTURE RACES
//
// "name":"Spriggan",
// "description":"[WIP]",
// "id":"SPRIGGAN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Gnome",
// "description":"[WIP]",
// "id":"GNOME",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Catlike",
// "description":"[WIP]",
// "id":"CATLIKE",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Dragonborn",
// "description":"[WIP]",
// "id":"DRAGONBORN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Minotaur",
// "description":"[WIP]",
// "id":"MINOTAUR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Orc",
// "description":"[WIP]",
// "id":"ORC",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Orc",
// "description":"[WIP]",
// "id":"HALFORC",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Giant",
// "description":"[WIP]",
// "id":"HALFGIANT",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Goblin",
// "description":"[WIP]",
// "id":"GOBLIN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Centaur",
// "description":"[WIP]",
// "id":"CENTAUR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Angel",
// "description":"[WIP]",
// "id":"HALFANGEL",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Tiefling",
// "description":"[WIP]",
// "id":"TIEFLING",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Elf",
// "description":"[WIP]",
// "id":"HALFELF",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Automaton",
// "description":"[WIP]",
// "id":"AUTOMATON",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Satyr",
// "description":"[WIP]",
// "id":"SATYR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Eutara",
// "description":"Hospitable ocean-dwellers",
// "id":"EUTARA",
// "long_description": "<p>A lot more about dwarves mother fuckers. We can put <strong>HTML</strong> in here.</p>"
//
// "name":"Halfling",
// "description":"Peaceful and practical. Halflings ",
// "id":"HALFLING",
// "long_description": "<p>This is a work in progress.</p>"
