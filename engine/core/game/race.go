package game

type Race struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	LongDescription string `json:"long_description"`
	ID              string `json:"id, omitempty"`
}