package core

type Gender struct {
	Name            string   `json:"name"`
	ID              string   `json:"id"`
	DisallowedRaces []string `json:"disallowed_races"`
	Only            string   `json:"only"`
	Human           string   `json:"-"`
	Scientific      string   `json:"-"`
	Pronoun         string   `json:"-"`
}

func (g *Gender) RaceAllowed(race string) bool {
	if g.Only != "" && g.Only != race { // only value exists and doesn't match provided race
		return false
	}

	for _, value := range g.DisallowedRaces {
		if value == race { // race is in the disallowed list
			return false
		}
	}

	return true
}
