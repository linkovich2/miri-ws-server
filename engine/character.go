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
		// check for current step and handle accordingly
		switch form.Step {

		}
	} else {
		activeCharacterForms[u] = &CharacterForm{Step: CharCreate_Start}
		hub.BasicSend("charcreate", races, u.Connection)
	}
}

func (h *HandlerInterface) CommandAuthenticated_CHARCREATE_STEPBACK(u *User, args *json.RawMessage) {
	// @todo
}
