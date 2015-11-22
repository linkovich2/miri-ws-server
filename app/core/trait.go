package core

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
