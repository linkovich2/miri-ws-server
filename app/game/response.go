package game

type (
	ResponseLocation struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	response struct {
		Location ResponseLocation `json:"location"`
		Messages []string         `json:"messages"`
	}
)
