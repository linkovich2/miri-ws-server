package core

import (
	"reflect"
	"strings"
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
	room.Message("<default>You take a drink from the water well.</default>", initiator, callback)
	room.BroadcastToAllButCharacter(
		strings.Join([]string{"<default>", initiator.ShortDescriptionWithName(), " takes a drink from the water well.</default>"}, ""),
		initiator,
		callback,
	)
}

func (i DrinkFromInteraction) Title() string {
	return "Drink"
}
