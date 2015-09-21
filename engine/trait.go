package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type (
	TraitCategory struct {
		Name   string `json:"name"`
		Unique bool   `json:"unique"`
		ID     string `json:"id"`
	}

	AestheticTrait struct {
		Name              string   `json:"name"`
		ID                string   `json:"id"`
		Description       string   `json:"description"`
		Image             int      `json:"image"`
		Category          string   `json:"category"`
		DisallowedRaces   []string `json:"disallowed_races"`
		DisallowedGenders []string `json:"disallowed_genders"`
	}

	AestheticTraitShort struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       int    `json:"image"`
		Category    string `json:"category"`
		Unique      bool   `json:"unique"`
	}

	FunctionalTrait      struct{} // @todo
	FunctionalTraitShort struct{}
)

var (
	aestheticTraits  = make(map[string]AestheticTrait)
	functionalTraits = make(map[string]FunctionalTrait)
	traitCategories  = make(map[string]TraitCategory)
)

func InitAestheticTraits() {
	arr := []AestheticTrait{}
	loader.Grab("aesthetic_traits.json", &arr)
	for _, val := range arr {
		aestheticTraits[val.ID] = val
	}
}

func InitTraitCategories() {
	arr := []TraitCategory{}
	loader.Grab("trait_categories.json", &arr)
	for _, val := range arr {
		traitCategories[val.ID] = val
	}
}

func (a *AestheticTrait) Shorten() (short AestheticTraitShort) {
	short = AestheticTraitShort{
		Name:        a.Name,
		ID:          a.ID,
		Description: a.Description,
		Image:       a.Image,
		Category:    a.Category,
		Unique:      a.Unique,
	}

	return short
}
