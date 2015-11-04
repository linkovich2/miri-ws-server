package models

type Character struct {
	Race             string              `json:"race"`
	Gender           string              `json:"gender"`
	AestheticTraits  map[string][]string `json:"aesthetic_traits"`
	FunctionalTraits map[string][]string `json:"functional_traits"`
	Background       string              `json:"background"`
	Name             string              `json:"name"`
}

// map[string][]string in JSON might be { "HAIRCOLOR": [ "BLONDE" ], "OTHERAESTHETICS": [ "FRECKLEY","REDNOSED" ] }
// This way we can access traits by going AestheticTraitCategories[cat].Traits[trait].Description etc
