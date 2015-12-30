package core

type (
	ComponentBag struct {
		Name         string             `json:"name"`
		Interactions []Interaction      `json:"interactions"`
		Properties   PropertyCollection `json:"-"`
		Behaviors    []Behavior         `json:"-"`
		Visible      bool               `json:"visible"`
	}
	PropertyCollection []*Property
	Property           struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	Interaction struct {
		Command string `json:"command"`
		// @todo
	}

	Entity interface {
		Update(*Room, func(string, string))
	}
)

func (c *ComponentBag) Update(r *Room, callback func(string, string)) {
	for _, b := range c.Behaviors {
		b.Perform(c, r, callback)
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
