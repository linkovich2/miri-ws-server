package engine

import ()

type Gender struct {
	Name            string   `json:"name"`
	ID              string   `json:"id"`
	DisallowedRaces []string `json:"disallowed_races"`
	Only            string   `json:"only"`
}

type GenderShort struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var genders = make(map[string]Gender)
