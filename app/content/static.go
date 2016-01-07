package content

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var (
	backgrounds      map[string]core.Background
	aestheticTraits  map[string]core.AestheticTraitCategory
	functionalTraits map[string]core.FunctionalTraitCategory
	races            map[string]core.Race
	genders          map[string]core.Gender
)

func AestheticTraits() map[string]core.AestheticTraitCategory {
	if len(aestheticTraits) <= 0 {
		data, err := Asset("json/aesthetic_traits.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.AestheticTraitCategory{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		aestheticTraits = a
	}

	return aestheticTraits
}

func AestheticTrait(key string) core.AestheticTraitCategory {
	t := AestheticTraits()
	return t[key]
}

func Backgrounds() map[string]core.Background {
	if len(backgrounds) <= 0 {
		data, err := Asset("json/backgrounds.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.Background{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		backgrounds = a
	}

	return backgrounds
}

func Background(key string) core.Background {
	t := Backgrounds()
	return t[key]
}

func FunctionalTraits() map[string]core.FunctionalTraitCategory {
	if len(functionalTraits) <= 0 {
		data, err := Asset("json/functional_traits.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.FunctionalTraitCategory{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		functionalTraits = a
	}

	return functionalTraits
}

func FunctionalTraitCategory(key string) core.FunctionalTraitCategory {
	t := FunctionalTraits()
	return t[key]
}

func Races() map[string]core.Race {
	if len(races) <= 0 {
		data, err := Asset("json/races.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.Race{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		races = a
	}

	return races
}

func Race(key string) core.Race {
	races := Races()
	return races[key]
}

func Genders() map[string]core.Gender {
	if len(genders) <= 0 {
		data, err := Asset("json/genders.json")
		if err != nil {
			panic(err)
		}

		a := map[string]core.Gender{}
		err = json.Unmarshal(data, &a)
		if err != nil {
			panic(err)
		}
		genders = a
	}

	return genders
}

func Gender(key string) core.Gender {
	genders := Genders()
	return genders[key]
}
