package core

import (
	"encoding/json"
)

var (
	aestheticTraits  map[string]AestheticTraitCategory
	functionalTraits map[string]FunctionalTraitCategory
)

type (
	AestheticTraitCategory struct {
		Name              string                    `json:"name"`
		Unique            bool                      `json:"unique"`
		ID                string                    `json:"id"`
		Traits            map[string]AestheticTrait `json:"traits"`
		Only              string                    `json:"only"`
		DisallowedRaces   []string                  `json:"disallowed_races"`
		DisallowedGenders []string                  `json:"disallowed_genders"`
		Minimum           int                       `json:"min"`
		Description       string                    `json:"description"`
	}

	FunctionalTraitCategory struct {
		Name              string                     `json:"name"`
		Unique            bool                       `json:"unique"`
		ID                string                     `json:"id"`
		Traits            map[string]FunctionalTrait `json:"traits"`
		Only              string                     `json:"only"`
		DisallowedRaces   []string                   `json:"disallowed_races"`
		DisallowedGenders []string                   `json:"disallowed_genders"`
		Minimum           int                        `json:"min"`
		Description       string                     `json:"description"`
	}

	AestheticTrait struct {
		Name              string   `json:"name"`
		ID                string   `json:"id"`
		Description       string   `json:"description"`
		Category          string   `json:"category"`
		DisallowedRaces   []string `json:"disallowed_races"`
		DisallowedGenders []string `json:"disallowed_genders"`
		Only              string   `json:"only"`
	}

	FunctionalTrait struct {
		Name              string   `json:"name"`
		ID                string   `json:"id"`
		Description       string   `json:"description"`
		Category          string   `json:"category"`
		DisallowedRaces   []string `json:"disallowed_races"`
		DisallowedGenders []string `json:"disallowed_genders"`
		Only              string   `json:"only"`
		Points            string   `json:"points"`
		Required          bool     `json:"required"`
	}
)

func (c *AestheticTraitCategory) IsAllowedForCharacter(character *Character) bool {
	return validateCharacterAgainst(character, c.DisallowedRaces, c.DisallowedGenders, c.Only)
}

func (t *AestheticTrait) IsAllowedForCharacter(character *Character) bool {
	return validateCharacterAgainst(character, t.DisallowedRaces, t.DisallowedGenders, t.Only)
}

func (c *FunctionalTraitCategory) IsAllowedForCharacter(character *Character) bool {
	return validateCharacterAgainst(character, c.DisallowedRaces, c.DisallowedGenders, c.Only)
}

func (t *FunctionalTrait) IsAllowedForCharacter(character *Character) bool {
	return validateCharacterAgainst(character, t.DisallowedRaces, t.DisallowedGenders, t.Only)
}

func validateCharacterAgainst(character *Character, disallowedRaces, disallowedGenders []string, only string) bool {
	if only != "" && character.Race != only && character.Gender != only {
		return false
	}

	for _, race := range disallowedRaces {
		if race == character.Race {
			return false
		}
	}

	for _, gender := range disallowedGenders {
		if gender == character.Gender {
			return false
		}
	}

	return true
}

// Content methods
func GetAestheticTraits() map[string]AestheticTraitCategory {
	if len(aestheticTraits) <= 0 {
		data, err := Asset("json/aesthetic_traits.json")
		if err != nil {
			panic(err)
		}

		a := map[string]AestheticTraitCategory{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		aestheticTraits = a
	}

	return aestheticTraits
}

func GetAestheticTrait(key string) AestheticTraitCategory {
	t := GetAestheticTraits()
	return t[key]
}

func GetFunctionalTraits() map[string]FunctionalTraitCategory {
	if len(functionalTraits) <= 0 {
		data, err := Asset("json/functional_traits.json")
		if err != nil {
			panic(err)
		}

		a := map[string]FunctionalTraitCategory{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		functionalTraits = a
	}

	return functionalTraits
}

func GetFunctionalTraitCategory(key string) FunctionalTraitCategory {
	t := GetFunctionalTraits()
	return t[key]
}
