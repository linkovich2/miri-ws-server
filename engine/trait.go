package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util/loader"
)

type (
	AestheticTraitCategory struct {
		Name              string           `json:"name"`
		Unique            bool             `json:"unique"`
		ID                string           `json:"id"`
		Traits            []AestheticTrait `json:"traits"`
		Only              string           `json:"only"`
		DisallowedRaces   []string         `json:"disallowed_races"`
		DisallowedGenders []string         `json:"disallowed_genders"`
		Minimum           int              `json:"min"`
	}

	AestheticTraitCategoryShort struct {
		Name    string                `json:"name"`
		Unique  bool                  `json:"unique"`
		ID      string                `json:"id"`
		Traits  []AestheticTraitShort `json:"traits"`
		Minimum int                   `json:"min"`
	}

	FunctionalTraitCategory struct {
		Name              string            `json:"name"`
		Unique            bool              `json:"unique"`
		ID                string            `json:"id"`
		Traits            []FunctionalTrait `json:"traits"`
		Only              string            `json:"only"`
		DisallowedRaces   []string          `json:"disallowed_races"`
		DisallowedGenders []string          `json:"disallowed_genders"`
		Minimum           int               `json:"min"`
	}

	FunctionalTraitCategoryShort struct {
		Name    string                 `json:"name"`
		Unique  bool                   `json:"unique"`
		ID      string                 `json:"id"`
		Traits  []FunctionalTraitShort `json:"traits"`
		Minimum int                    `json:"min"`
	}

	AestheticTrait struct {
		Name              string   `json:"name"`
		ID                string   `json:"id"`
		Description       string   `json:"description"`
		Image             int      `json:"image"`
		Category          string   `json:"category"`
		DisallowedRaces   []string `json:"disallowed_races"`
		DisallowedGenders []string `json:"disallowed_genders"`
		Only              string   `json:"only"`
	}

	AestheticTraitShort struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       int    `json:"image"`
	}

	FunctionalTrait struct {
		Name              string   `json:"name"`
		ID                string   `json:"id"`
		Description       string   `json:"description"`
		Image             int      `json:"image"`
		Category          string   `json:"category"`
		DisallowedRaces   []string `json:"disallowed_races"`
		DisallowedGenders []string `json:"disallowed_genders"`
		Only              string   `json:"only"`
		Points            string   `json:"points"`
	}

	FunctionalTraitShort struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       int    `json:"image"`
		Points      string `json:"points"`
	}
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
			t.Category = val.ID
			aestheticTraits[t.ID] = t
		}
	}
}

func InitFunctionalTraits() {
	arr := []FunctionalTraitCategory{}
	loader.Grab("functional_traits.json", &arr)
	for _, val := range arr {
		functionalTraitsCategorized[val.ID] = val
		for _, t := range val.Traits {
			t.Category = val.ID
			functionalTraits[t.ID] = t
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
