package engine

import (
	"encoding/json"
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
		Race   string `json:"race"`
		Gender string `json:"gender"`
		// AestheticTraits []AestheticTrait
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
			hub.BasicSend("charcreateaesthetic", nil, u.Connection)
		case CharCreate_Aesthetic:
			hub.BasicSend("charcreatefunctional", nil, u.Connection)
		case CharCreate_Functional:
			hub.BasicSend("charcreatebackgrounds", nil, u.Connection)
		case CharCreate_Background:
			hub.BasicSend("charcreatename", nil, u.Connection)
		case CharCreate_Name:
			// DONE! yay! save the character if everything is valid
		}

	} else {
		activeCharacterForms[u] = &CharacterForm{Step: CharCreate_Start}
		hub.BasicSend("charcreateraces", races, u.Connection)
	}
}

func (h *HandlerInterface) CommandAuthenticated_CHARCREATESTEPBACK(u *User, args *json.RawMessage) {
	// @todo
}

func (h *HandlerInterface) CommandAuthenticated_NEWCHAR(u *User, args *json.RawMessage) {
	if _, exists := activeCharacterForms[u]; exists {
		delete(activeCharacterForms, u)
	} else {
		logger.Error(" -- tried to cancel CHARCREATE form that didn't exist, ignoring.")
	}
}

func (f *CharacterForm) validateRace(c *Character) bool {
	valid := false

	for _, r := range races {
		if r.ID == c.Race {
			valid = true
			break
		}
	}

	return valid
}

func (f *CharacterForm) getAvailableGenders() []GenderShort {
	res := []GenderShort{}

	for _, g := range genders {
		remove := false

		for _, exclude := range g.DisallowedRaces {
			if exclude == f.Character.Race {
				remove = true
			}
		}

		if !remove {
			res = append(res, GenderShort{g.Name, g.ID})
		}
	}

	return res
}
