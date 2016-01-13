package core

import (
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"reflect"
)

var (
	InteractorRegistry = map[string]reflect.Type{
		"SPEAK_TO": reflect.TypeOf(SpeakToInteraction{}),
	}
)

type (
	Interactor interface {
		Perform(*ComponentBag, *Character, *Room, func(string, string))
		Title() string
	}

	SpeakToInteraction struct{}
)

func (i SpeakToInteraction) Perform(target *ComponentBag, initiator *Character, room *Room, callback func(string, string)) {
	logger.Write.Info("Character [%s] spoke to [%s]", initiator.Name, target.Name)
}

func (i SpeakToInteraction) Title() string {
	return "Speak"
}
