package core

import (
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var BehaviorRegistry = map[string]reflect.Type{
	"CROWD_CONVERSATION": reflect.TypeOf(CrowdConversationBehavior{}),
	"AMBIANCE":           reflect.TypeOf(AmbianceBehavior{}),
}

type (
	Behavior interface {
		Perform(*ComponentBag, *Room, func(string, string))
	}

	CrowdConversationBehavior struct{}
	AmbianceBehavior          struct{}
)

// @todo organize this into an init method and load up behaviors into the registry

func (c CrowdConversationBehavior) Perform(cb *ComponentBag, room *Room, callback func(string, string)) {
	conversations := [][]string{}

	state := cb.Properties.ValueOf("state")
	if state == "" {
		// no state found, let's decide whether we want to start a conversation
		rand.Seed(time.Now().UnixNano())
		// if we choose to start a conversation, we should add a "placeholder" prop and a "state" prop
		// if we are on the last part of the conversation, we should broadcast that and remove the state and placeholder props
		// from the bag
		freq := getFrequencyProp(cb)

		if rand.Intn(freq)+1 == freq {
			for _, p := range cb.Properties.Matching("conversation") {
				conversations = append(conversations, strings.Split(p.Value, ";;"))
			}
			if len(conversations) <= 0 { // failsafe
				return
			}

			// now we have conversations, let's pick one and decide whether we want to start a conversation now
			i := rand.Intn(len(conversations))
			cb.Properties = append(cb.Properties, &Property{Key: "state", Value: "conversing"})
			cb.Properties = append(cb.Properties, &Property{Key: "placeholder", Value: strconv.Itoa(i) + ";1"})
			room.Broadcast(conversations[i][0], callback)
		}
	} else if state == "conversing" {
		// we're conversing, we're gonna need to check out the placeholder prop and continue from where we left off
		for _, p := range cb.Properties.Matching("conversation") {
			conversations = append(conversations, strings.Split(p.Value, ";;"))
		}

		placeholder := cb.Properties.ValueOf("placeholder")
		splitPlaceholder := strings.Split(placeholder, ";")

		conversationIndex, _ := strconv.Atoi(splitPlaceholder[0])
		placeholderIndex, _ := strconv.Atoi(splitPlaceholder[1])
		convo := conversations[conversationIndex]
		if len(convo) < placeholderIndex+1 {
			cb.Properties = cb.Properties.Remove("placeholder")
			cb.Properties = cb.Properties.Remove("state")
			return
		}

		if convo[placeholderIndex] != "PAUSE" {
			room.Broadcast(convo[placeholderIndex], callback)
		}

		cb.Properties.Update("placeholder", strings.Join([]string{splitPlaceholder[0], strconv.Itoa(placeholderIndex + 1)}, ";"))
	}
}

func (a AmbianceBehavior) Perform(cb *ComponentBag, room *Room, callback func(string, string)) {
	rand.Seed(time.Now().UnixNano())
	freq := getFrequencyProp(cb)

	ambiance := []string{}
	if rand.Intn(freq)+1 == freq {
		for _, p := range cb.Properties.Matching("ambiance") {
			ambiance = append(ambiance, p.Value)
		}
		if len(ambiance) <= 0 { // failsafe
			return
		}

		room.Broadcast(ambiance[rand.Intn(len(ambiance))], callback)
	}
}

func getFrequencyProp(cb *ComponentBag) int {
	freqString := cb.Properties.ValueOf("frequency")
	if freqString == "" {
		return 10
	} else {
		c, _ := strconv.Atoi(freqString)
		return c
	}
}
