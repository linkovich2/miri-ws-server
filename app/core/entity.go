package core

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"reflect"
)

var e entityContainer

type (
	ComponentBag struct {
		Name         string             `json:"name"`
		Interactions []Interaction      `json:"interactions"`
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

	entityContainer map[string]*ComponentBag

	readEntity struct {
		Name         string     `json:"name"`
		Behaviors    []string   `json:"behaviors"`
		Properties   [][]string `json:"properties"`
		NotVisible   bool       `json:"not_visible"`
		Interactions []string   `json:"interactions"`
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

func init() {
	if len(e) <= 0 {
		e = entityContainer{}
		files, err := AssetDir("json/entities")
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			data, _ := Asset("json/entities/" + f)
			a := map[string]readEntity{}
			err := json.Unmarshal(data, &a)
			if err != nil {
				panic(err)
			}

			for key, re := range a {
				c := &ComponentBag{
					Name:       re.Name,
					NotVisible: re.NotVisible,
				}

				properties := PropertyCollection{}
				for _, p := range re.Properties {
					properties = append(properties, &Property{p[0], p[1]})
				}
				c.Properties = properties

				behaviors := []Behavior{}
				for _, b := range re.Behaviors {
					behaviors = append(
						behaviors,
						reflect.New(BehaviorRegistry[b]).Elem().Interface().(Behavior),
					)
				}
				c.Behaviors = behaviors

				interactions := []Interaction{}
				for _, i := range re.Interactions {
					interactor := reflect.New(InteractorRegistry[i]).Elem().Interface().(Interactor)
					interactions = append(
						interactions,
						Interaction{interactor.Title(), interactor},
					)
				}
				c.Interactions = interactions

				e[key] = c
			}
		}
	}
}

func (ec entityContainer) get(i string) *ComponentBag {
	if entity, exists := ec[i]; exists {
		return entity.Copy()
	} else {
		panic("Entity doesn't exist!") // @todo make this more specific
	}
}
