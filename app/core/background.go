package core

import "encoding/json"

var backgrounds map[string]Background

type (
	Background struct {
		Name          string                  `json:"name"`
		ID            string                  `json:"id"`
		Prerequisites BackgroundPrerequisites `json:"prerequisites"`
		AllowAll      bool                    `json:"allow_all"`
		Description   string                  `json:"description"`
		StartPosition string                  `json:"start_position"`
		StartRealm    string                  `json:"start_realm"`
	}

	BackgroundPrerequisites struct {
		Races            []string `json:"races"`
		Genders          []string `json:"genders"`
		AestheticTraits  []string `json:"aesthetic_traits"`
		FunctionalTraits []string `json:"functional_traits"`
	}
)

func (b *Background) IsAllowedForCharacter(character *Character) bool {
	if b.AllowAll {
		return true
	}

	matchesRace, matchesGender := true, true

	if len(b.Prerequisites.Races) > 0 {
		matchesRace = false
		for _, r := range b.Prerequisites.Races {
			if character.Race == r {
				matchesRace = true
				break
			}
		}
	}

	if len(b.Prerequisites.Genders) > 0 {
		matchesGender = false
		for _, g := range b.Prerequisites.Genders {
			if character.Gender == g {
				matchesGender = true
				break
			}
		}
	}

	if len(b.Prerequisites.AestheticTraits) > 0 {
		var traitMatches = make(map[string]bool)
		for _, a := range b.Prerequisites.AestheticTraits {
			traitMatches[a] = false
			for _, c := range character.AestheticTraits {
				for _, t := range c {
					if t == a {
						traitMatches[a] = true
						break
					}
				}
			}
		}

		for _, match := range traitMatches {
			if !match {
				return false
			}
		}
	}

	if len(b.Prerequisites.FunctionalTraits) > 0 {
		var traitMatches = make(map[string]bool)
		for _, a := range b.Prerequisites.FunctionalTraits {
			traitMatches[a] = false
			for _, c := range character.FunctionalTraits {
				for _, t := range c {
					if t == a {
						traitMatches[a] = true
						break
					}
				}
			}
		}

		for _, match := range traitMatches {
			if !match {
				return false
			}
		}
	}

	return (matchesRace && matchesGender)
}

// Content methods
func GetBackgrounds() map[string]Background {
	if len(backgrounds) <= 0 {
		data, err := Asset("json/backgrounds.json")
		if err != nil {
			panic(err)
		}

		a := map[string]Background{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		backgrounds = a
	}

	return backgrounds
}

func GetBackground(key string) Background {
	t := GetBackgrounds()
	return t[key]
}
