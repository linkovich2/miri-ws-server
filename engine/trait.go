package engine

import (
	"github.com/jonathonharrell/miri-ws-server/engine/util"
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
		Required          bool     `json:"required"`
	}

	FunctionalTraitShort struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       int    `json:"image"`
		Points      string `json:"points"`
		Required    bool   `json:"required"`
	}
)

var (
	aestheticTraits             = make(map[string]AestheticTrait)
	functionalTraits            = make(map[string]FunctionalTrait)
	aestheticTraitsCategorized  = make(map[string]AestheticTraitCategory)
	functionalTraitsCategorized = make(map[string]FunctionalTraitCategory)
)

func (ac *AestheticTraitCategory) AvailableForCharacter(c *Character) bool {
	if in, _ := util.InArray(c.Race, ac.DisallowedRaces); in {
		return false
	}

	if in, _ := util.InArray(c.Gender, ac.DisallowedGenders); in {
		return false
	}

	if ac.Only != "" && c.Race != ac.Only && c.Gender != ac.Only {
		return false
	}

	return true
}

func (a *AestheticTrait) Shorten() (short AestheticTraitShort) {
	short = AestheticTraitShort{
		Name:        a.Name,
		ID:          a.ID,
		Description: a.Description,
		Image:       a.Image,
	}

	return
}

func (f *FunctionalTrait) Shorten() (short FunctionalTraitShort) {
	short = FunctionalTraitShort{
		Name:        f.Name,
		ID:          f.ID,
		Description: f.Description,
		Image:       f.Image,
		Points:      f.Points,
		Required:    f.Required,
	}

	return
}
