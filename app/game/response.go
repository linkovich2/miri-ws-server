package game

type (
	ResponseLocation struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Messages    []string `json:"messages"`
	}

	response struct {
		Location ResponseLocation `json:"location"`
	}
)
