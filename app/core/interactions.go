package core

import (
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"reflect"
)

var (
	InteractorRegistry = map[string]reflect.Type{
		"DRINK_FROM": reflect.TypeOf(DrinkFromInteraction{}),
	}
)

type (
	Interactor interface {
		Perform(*ComponentBag, *Character, *Room, func(string, string))
		Title() string
	}

	DrinkFromInteraction struct{}
)

func (i DrinkFromInteraction) Perform(target *ComponentBag, initiator *Character, room *Room, callback func(string, string)) {
	room.Message("You take a drink from the water well.", initiator, callback)
	// @todo this needs to broadcast something else to the rest of the people in the room
	logger.Write.Info("Character [%s] tried to drink from [%s]", initiator.Name, target.Name)
}

func (i DrinkFromInteraction) Title() string {
	return "Drink"
}
