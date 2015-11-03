package models

type Character struct {
	Race             string   `json:"race"`
	Gender           string   `json:"gender"`
	AestheticTraits  []string `json:"aesthetic_traits"`
	FunctionalTraits []string `json:"functional_traits"`
	Background       string   `json:"background"`
	Name             string   `json:"name"`
}
