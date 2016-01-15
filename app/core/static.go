package core

import (
	"encoding/json"
)

var (
	backgrounds      map[string]Background
	aestheticTraits  map[string]AestheticTraitCategory
	functionalTraits map[string]FunctionalTraitCategory
	races            map[string]Race
	genders          map[string]Gender
)

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

func GetGenders() map[string]Gender {
	if len(genders) <= 0 {
		data, err := Asset("json/genders.json")
		if err != nil {
			panic(err)
		}

		a := map[string]Gender{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		genders = a
	}

	return genders
}

func GetGender(key string) Gender {
	genders := GetGenders()
	return genders[key]
}
