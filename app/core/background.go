package core

type (
	Background struct {
		Name          string                  `json:"name"`
		ID            string                  `json:"id"`
		Prerequisites BackgroundPrerequisites `json:"prerequisites"`
		AllowAll      bool                    `json:"allow_all"`
		Description   string                  `json:"description"`
	}

	BackgroundPrerequisites struct {
		Races            []string
		Genders          []string
		AestheticTraits  []string
		FunctionalTraits []string
	}
)

func (b *Background) IsAllowedForCharacter(character *Character) bool {
	if b.AllowAll {
		return true
	}

	matchesRace, matchesGender := true, true

	if len(b.Prerequisites.Races) > 0 {
		matchesRace = false
		for _, r := range b.Prerequisites.Races {
			if character.Race == r {
				matchesRace = true
				break
			}
		}
	}

	if len(b.Prerequisites.Genders) > 0 {
		matchesGender = false
		for _, g := range b.Prerequisites.Genders {
			if character.Gender == g {
				matchesGender = true
				break
			}
		}
	}

	if len(b.Prerequisites.AestheticTraits) > 0 {
		var traitMatches = make(map[string]bool)
		for _, a := range b.Prerequisites.AestheticTraits {
			traitMatches[a] = false
			for _, c := range character.AestheticTraits {
				for _, t := range c {
					if t == a {
						traitMatches[a] = true
						break
					}
				}
			}
		}

		for _, match := range traitMatches {
			if !match {
				return false
			}
		}
	}

	if len(b.Prerequisites.FunctionalTraits) > 0 {
		var traitMatches = make(map[string]bool)
		for _, a := range b.Prerequisites.FunctionalTraits {
			traitMatches[a] = false
			for _, c := range character.FunctionalTraits {
				for _, t := range c {
					if t == a {
						traitMatches[a] = true
						break
					}
				}
			}
		}

		for _, match := range traitMatches {
			if !match {
				return false
			}
		}
	}

	return (matchesRace && matchesGender)
}
