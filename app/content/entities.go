package content

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"reflect"
)

type entityContainer map[string]*core.ComponentBag

var e entityContainer

type readEntity struct {
	Name         string     `json:"name"`
	Behaviors    []string   `json:"behaviors"`
	Properties   [][]string `json:"properties"`
	NotVisible   bool       `json:"not_visible"`
	Interactions []string   `json:"interactions"`
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
				c := &core.ComponentBag{
					Name:       re.Name,
					NotVisible: re.NotVisible,
				}

				properties := core.PropertyCollection{}
				for _, p := range re.Properties {
					properties = append(properties, &core.Property{p[0], p[1]})
				}
				c.Properties = properties

				behaviors := []core.Behavior{}
				for _, b := range re.Behaviors {
					behaviors = append(
						behaviors,
						reflect.New(core.BehaviorRegistry[b]).Elem().Interface().(core.Behavior),
					)
				}
				c.Behaviors = behaviors

				interactions := []core.Interaction{}
				for _, i := range re.Interactions {
					interactor := reflect.New(core.InteractorRegistry[i]).Elem().Interface().(core.Interactor)
					interactions = append(
						interactions,
						core.Interaction{interactor.Title(), interactor},
					)
				}
				c.Interactions = interactions

				e[key] = c
			}
		}
	}
}

func (ec entityContainer) get(i string) *core.ComponentBag {
	if entity, exists := ec[i]; exists {
		return entity.Copy()
	} else {
		panic("Entity doesn't exist!") // @todo make this more specific
	}
}
