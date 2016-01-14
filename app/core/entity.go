package core

import "github.com/jonathonharrell/miri-ws-server/app/logger"

type (
	ComponentBag struct {
		Name         string             `json:"name"`
		Interactions []Interaction      `json:"-"`
		Properties   PropertyCollection `json:"-"`
		Behaviors    []Behavior         `json:"-"`
		NotVisible   bool               `json:"not_visible"`
	}
	PropertyCollection []*Property
	Property           struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	Interaction struct {
		Title  string     `json:"title"`
		Action Interactor `json:"-"`
	}

	Entity interface {
		Update(*Room, func(string, string))
		Interact(string, *Character, *Room, func(string, string))
	}
)

func (c *ComponentBag) Update(r *Room, callback func(string, string)) {
	for _, b := range c.Behaviors {
		b.Perform(c, r, callback)
	}
}

func (c *ComponentBag) Interact(action string, character *Character, r *Room, callback func(string, string)) {
	for _, i := range c.Interactions {
		if i.Action.Title() == action {
			i.Action.Perform(c, character, r, callback)
			return
		}
	}

	logger.Write.Error("Character [%s] sent an action [%s] that is not available on this entity [%s].", character.Name, action, c.Name)
}

func (c *ComponentBag) Copy() *ComponentBag {
	return &ComponentBag{
		Name:         c.Name,
		Interactions: c.Interactions,
		Properties:   c.Properties,
		Behaviors:    c.Behaviors,
		NotVisible:   c.NotVisible,
	}
}

func (pc PropertyCollection) ValueOf(key string) string {
	for _, p := range pc {
		if p.Key == key {
			return p.Value
		}
	}

	return ""
}

func (pc PropertyCollection) Matching(key string) []*Property {
	res := []*Property{}
	for _, p := range pc {
		if p.Key == key {
			res = append(res, p)
		}
	}

	return res
}

func (pc PropertyCollection) Update(key, value string) {
	for _, p := range pc {
		if p.Key == key {
			p.Value = value
		}
	}
}

func (pc PropertyCollection) Remove(key string) PropertyCollection {
	for i, p := range pc {
		if p.Key == key {
			pc = append(pc[:i], pc[i+1:]...)
			break
		}
	}

	return pc
}
