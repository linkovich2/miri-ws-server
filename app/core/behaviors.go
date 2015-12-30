package core

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type (
	Behavior interface {
		Perform(*ComponentBag, *Room, func(string, string))
	}

	CrowdConversationBehavior struct{}
	Ambiance                  struct{}
)

func (c CrowdConversationBehavior) Perform(cb *ComponentBag, room *Room, callback func(string, string)) {
	conversations := [][]string{}

	state := cb.Properties.ValueOf("state")
	if state == "" {
		// no state found, let's decide whether we want to start a conversation
		rand.Seed(time.Now().UnixNano())
		// if we choose to start a conversation, we should add a "placeholder" prop and a "state" prop
		// if we are on the last part of the conversation, we should broadcast that and remove the state and placeholder props
		// from the bag

		if rand.Intn(2)+1 == 2 {
			for _, p := range cb.Properties.Matching("conversation") {
				conversations = append(conversations, strings.Split(p.Value, ";;"))
			}
			if len(conversations) <= 0 { // failsafe
				return
			}

			// now we have conversations, let's pick one and decide whether we want to start a conversation now
			i := rand.Intn(len(conversations))
			cb.Properties = append(cb.Properties, &Property{Key: "state", Value: "conversing"})
			cb.Properties = append(cb.Properties, &Property{Key: "placeholder", Value: strconv.Itoa(i) + ";1;skip"})
			room.Broadcast(conversations[i][0], callback)
		}
	} else if state == "conversing" {
		// we're conversing, we're gonna need to check out the placeholder prop and continue from where we left off
		for _, p := range cb.Properties.Matching("conversation") {
			conversations = append(conversations, strings.Split(p.Value, ";;"))
		}

		placeholder := cb.Properties.ValueOf("placeholder")
		splitPlaceholder := strings.Split(placeholder, ";")
		if len(splitPlaceholder) > 2 {
			cb.Properties.Update("placeholder", strings.Replace(placeholder, ";skip", "", -1))
			return
		}

		conversationIndex, _ := strconv.Atoi(splitPlaceholder[0])
		placeholderIndex, _ := strconv.Atoi(splitPlaceholder[1])
		convo := conversations[conversationIndex]
		if len(convo) < placeholderIndex+1 {
			cb.Properties = cb.Properties.Remove("placeholder")
			cb.Properties = cb.Properties.Remove("state")
			return
		}
		room.Broadcast(convo[placeholderIndex], callback)

		cb.Properties.Update("placeholder", strings.Join([]string{splitPlaceholder[0], strconv.Itoa(placeholderIndex + 1), "skip"}, ";"))
	}
}

func (a Ambiance) Perform(cb *ComponentBag, room *Room, callback func(string, string)) {
	// @stub
}
