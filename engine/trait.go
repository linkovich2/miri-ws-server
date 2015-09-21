package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type (
	AestheticTraitCategory struct {
		Name   string           `json:"name"`
		Unique bool             `json:"unique"`
		ID     string           `json:"id"`
		Traits []AestheticTrait `json:"traits"`
	}

	AestheticTraitCategoryShort struct {
		Name   string                `json:"name"`
		Unique bool                  `json:"unique"`
		ID     string                `json:"id"`
		Traits []AestheticTraitShort `json:"traits"`
	}

	FunctionalTraitCategory struct {
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
	}

	FunctionalTrait      struct{} // @todo
	FunctionalTraitShort struct{}
)

var (
	aestheticTraits             = make(map[string]AestheticTrait)
	functionalTraits            = make(map[string]FunctionalTrait)
	aestheticTraitsCategorized  = make(map[string]AestheticTraitCategory)
	functionalTraitsCategorized = make(map[string]FunctionalTraitCategory)
)

func InitAestheticTraits() {
	arr := []AestheticTraitCategory{}
	loader.Grab("aesthetic_traits.json", &arr)
	for _, val := range arr {
		aestheticTraitsCategorized[val.ID] = val
		for _, t := range val.Traits {
			aestheticTraits[t.ID] = t
		}
	}
}

func (a *AestheticTrait) Shorten() (short AestheticTraitShort) {
	short = AestheticTraitShort{
		Name:        a.Name,
		ID:          a.ID,
		Description: a.Description,
		Image:       a.Image,
	}

	return short
}
