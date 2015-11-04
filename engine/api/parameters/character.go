package parameters

type Character struct {
	Race             string              `json:"race"`
	Gender           string              `json:"gender"`
	AestheticTraits  map[string][]string `json:"aesthetic_traits"`
	FunctionalTraits map[string][]string `json:"functional_traits"`
	Background       string
	Name             string
}
