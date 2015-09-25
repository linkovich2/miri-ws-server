package engine

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/engine/util"
)

const (
	CharCreate_Start = iota
	CharCreate_Gender
	CharCreate_Aesthetic
	CharCreate_Functional
	CharCreate_Background
	CharCreate_Name
)

type (
	Character struct {
		Race            string   `json:"race"`
		Gender          string   `json:"gender"`
		AestheticTraits []string `json:"aesthetic_traits"`
		// FunctionalTraits []FunctionalTrait
		// Background string
		// Name string
	}

	CharacterForm struct {
		Character Character
		Step      int
	}
)

var activeCharacterForms = make(map[*User]*CharacterForm)

func (h *HandlerInterface) CommandAuthenticated_CHARLIST(u *User, args *json.RawMessage) {
	hub.BasicSend("charlist", u.Account.Characters, u.Connection)
}

func (h *HandlerInterface) CommandAuthenticated_CHARCREATE(u *User, args *json.RawMessage) {
	if form, exists := activeCharacterForms[u]; exists {
		// logger.Info("Connection [%s] active form on CHARCREATE STEP %v", u.Connection.ID, form.Step) // @temp

		// check for current step and handle accordingly
		switch form.Step {
		case CharCreate_Start:
			c := &Character{}
			_ = json.Unmarshal(*args, c)

			if form.validateRace(c) {
				// validate then send back genders
				form.Step = CharCreate_Gender
				form.Character.Race = c.Race
				hub.BasicSend("charcreatestepup", nil, u.Connection)
				hub.BasicSend("charcreategenders", form.getAvailableGenders(), u.Connection)
			}

		case CharCreate_Gender:
			c := &Character{}
			_ = json.Unmarshal(*args, c)

			if form.validateGender(c) {
				// validate then send back genders
				form.Step = CharCreate_Aesthetic
				form.Character.Gender = c.Gender
				hub.BasicSend("charcreatestepup", nil, u.Connection)
				hub.BasicSend("charcreateaesthetic", form.getAvailableAestheticTraits(), u.Connection)
			}
		case CharCreate_Aesthetic:
			c := &Character{}
			_ = json.Unmarshal(*args, c)

			valid, errors := form.validateAestheticTraits(c)
			if valid {
				form.Step = CharCreate_Functional
				form.Character.AestheticTraits = c.AestheticTraits
				hub.BasicSend("charcreatestepup", nil, u.Connection)
				hub.BasicSend("charcreatefunctional", form.getAvailableFunctionalTraits(), u.Connection)
			} else {
				hub.Send(&MessageResponse{Errors: errors, Success: false, ResponseTo: "charcreatestepup", Data: nil}, u.Connection)
			}
		case CharCreate_Functional:
			hub.BasicSend("charcreatebackgrounds", nil, u.Connection)
		case CharCreate_Background:
			hub.BasicSend("charcreatename", nil, u.Connection)
		case CharCreate_Name:
			// DONE! yay! save the character if everything is valid
		}

	} else {
		activeCharacterForms[u] = &CharacterForm{Step: CharCreate_Start}
		hub.BasicSend("charcreateraces", activeCharacterForms[u].getAvailableRaces(), u.Connection)
	}
}

func (h *HandlerInterface) CommandAuthenticated_CHARCREATESTEPBACK(u *User, args *json.RawMessage) {
	if form, exists := activeCharacterForms[u]; exists {
		if form.Step >= 1 {
			form.stepBack()
			hub.BasicSend("charcreatestepback", nil, u.Connection)
		} else {
			logger.Warning(" -- received step back on step 0 CHARCREATE, weird.")
		}
	} else {
		logger.Warning(" -- tried to step back CHARCREATE form that didn't exist, ignoring.")
	}
}

func (h *HandlerInterface) CommandAuthenticated_NEWCHAR(u *User, args *json.RawMessage) {
	if _, exists := activeCharacterForms[u]; exists {
		delete(activeCharacterForms, u)
	} else {
		logger.Warning(" -- tried to cancel CHARCREATE form that didn't exist, ignoring.")
	}
}

func (f *CharacterForm) stepBack() {
	f.Step = f.Step - 1
	switch f.Step {
	case CharCreate_Start:
		f.Character = Character{}
	case CharCreate_Gender:
		f.Character = Character{
			Race: f.Character.Race,
		}
	case CharCreate_Aesthetic:
		f.Character = Character{
			Race:   f.Character.Race,
			Gender: f.Character.Gender,
		}
	case CharCreate_Functional:
		f.Character = Character{
			Race:            f.Character.Race,
			Gender:          f.Character.Gender,
			AestheticTraits: f.Character.AestheticTraits,
		}
	case CharCreate_Background:
		// @todo character has race, gender, aesthetic and functional traits
	}
}

func (f *CharacterForm) validateRace(c *Character) bool {
	if _, exists := races[c.Race]; exists {
		return true
	}

	return false
}

func (f *CharacterForm) validateGender(c *Character) bool {
	if f.validateRace(c) {
		if _, exists := genders[c.Gender]; exists {
			return true
		}
	}

	return false
}

func (f *CharacterForm) validateAestheticTraits(c *Character) (valid bool, errors []string) {
	errors = []string{"You must select at least one hair style."}
	return
}

func (f *CharacterForm) getAvailableRaces() []Race {
	res := []Race{}

	for _, r := range races {
		if r.ID == "HUMAN" { // explicitly put human first
			res = append([]Race{r}, res...)
		} else {
			res = append(res, r)
		}
	}

	return res
}

func (f *CharacterForm) getAvailableGenders() []GenderShort {
	res := []GenderShort{}

	for _, g := range genders {
		if exclude, _ := util.InArray(f.Character.Race, g.DisallowedRaces); exclude {
			continue
		}

		if g.Only != "" && f.Character.Race != g.Only {
			continue
		}

		res = append(res, GenderShort{g.Name, g.ID})
	}

	return res
}

func (f *CharacterForm) getAvailableAestheticTraits() map[string]AestheticTraitCategoryShort {
	res := make(map[string]AestheticTraitCategoryShort)

	for _, cat := range aestheticTraitsCategorized { // Category Level
		if cat.Only != "" && f.Character.Race != cat.Only && f.Character.Gender != cat.Only {
			continue
		}

		if exclude, _ := util.InArray(f.Character.Race, cat.DisallowedRaces); exclude {
			continue
		}

		if exclude, _ := util.InArray(f.Character.Gender, cat.DisallowedGenders); exclude {
			continue
		}

		list := []AestheticTraitShort{}

		for _, t := range cat.Traits { // Trait inner loop
			if t.Only != "" && f.Character.Race != t.Only && f.Character.Gender != t.Only {
				continue
			}

			if exclude, _ := util.InArray(f.Character.Gender, t.DisallowedGenders); exclude {
				continue
			}

			if exclude, _ := util.InArray(f.Character.Race, t.DisallowedRaces); exclude {
				continue
			}

			list = append(list, t.Shorten())
		}

		res[cat.ID] = AestheticTraitCategoryShort{cat.Name, cat.Unique, cat.ID, list, cat.Minimum}
	}

	return res
}

func (f *CharacterForm) getAvailableFunctionalTraits() map[string]FunctionalTraitCategoryShort {
	res := make(map[string]FunctionalTraitCategoryShort)
	return res
}
