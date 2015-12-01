package core

import (
	// "encoding/json"
	// "github.com/jonathonharrell/miri-ws-server/app/util"
	// "strconv"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Character struct {
	ID               bson.ObjectId       `bson:"_id,omitempty" json:"-"`
	Race             string              `json:"race"`
	Gender           string              `json:"gender"`
	AestheticTraits  map[string][]string `json:"aesthetic_traits"`
	FunctionalTraits map[string][]string `json:"functional_traits"`
	Background       string              `json:"background"`
	Name             string              `json:"name"`
	UserID           string              `json:"-" bson:"user_id"`
	Created          time.Time           `json:"created"`
}

//
// func (f *CharacterForm) validateAestheticTraits(c *Character) (valid bool, errors []string) {
// 	atLeastOneRequired := make(map[string]bool) // [categoryId]satisfied{false}, for at least one is required
// 	hasOneIn := make(map[string]bool)           // for uniqueness validation, do we already have one in this field
// 	for _, cat := range aestheticTraitsCategorized {
// 		if cat.Minimum > 0 && cat.AvailableForCharacter(c) {
// 			atLeastOneRequired[cat.ID] = false
// 		}
// 	}
//
// 	for _, traitId := range c.AestheticTraits {
// 		if _, exists := aestheticTraits[traitId]; !exists {
// 			return false, []string{"Trait not found: " + traitId}
// 		}
//
// 		alreadyHad := hasOneIn[aestheticTraitsCategorized[aestheticTraits[traitId].Category].ID]
// 		hasOneIn[aestheticTraitsCategorized[aestheticTraits[traitId].Category].ID] = true
// 		if aestheticTraitsCategorized[aestheticTraits[traitId].Category].Unique && alreadyHad {
// 			return false, []string{"You may only have one " + aestheticTraitsCategorized[aestheticTraits[traitId].Category].Name + " selected."}
// 		}
//
// 		atLeastOneRequired[aestheticTraits[traitId].Category] = true
//
// 		// @todo also need to check this trait is valid for this race / gender
// 	}
//
// 	for i, cat := range atLeastOneRequired {
// 		if !cat {
// 			return false, []string{"You need to select at least one " + aestheticTraitsCategorized[i].Name}
// 		}
// 	}
//
// 	return true, []string{}
// }
//
// func (f *CharacterForm) validateFunctionalTraits(c *Character) (valid bool, errors []string) {
// 	validTraits := f.getAvailableFunctionalTraits()
// 	points := 0
// 	list := []string{}
//
// 	for _, ftc := range validTraits {
// 		for _, ft := range ftc.Traits {
// 			list = append(list, ft.ID)
//
// 			if in, _ := util.InArray(ft.ID, c.FunctionalTraits); !in {
// 				if ft.Required { // it's not in but it's required
// 					return false, []string{"'" + ft.Name + "' is a non-optional trait."}
// 				}
// 			} else {
// 				// it is in, validate anything else against it
// 				val, _ := strconv.Atoi(ft.Points)
// 				points = points + val
//
// 				// @todo need to validate for atleastonerequired, and any other constraints
// 			}
// 		}
// 	}
//
// 	if points < 0 {
// 		return false, []string{"You must have at least 0 points to continue."}
// 	}
//
// 	for _, trait := range c.FunctionalTraits {
// 		if in, _ := util.InArray(trait, list); !in {
// 			return false, []string{"Illegal trait included."}
// 		}
// 	}
//
// 	return true, []string{}
// }
//
// func (f *CharacterForm) validateBackground(c *Character) bool {
// 	// @todo stub
// 	return false
// }
//
// func (f *CharacterForm) validateName(c *Character) bool {
// 	// @todo stub
// 	return false
// }
//
// func (f *CharacterForm) save() {
// 	// @todo stub
// 	return
// }
//
// func (f *CharacterForm) getAvailableRaces() []Race {
// 	res := []Race{}
//
// 	for _, r := range races {
// 		if r.ID == "HUMAN" { // explicitly put human first
// 			res = append([]Race{r}, res...)
// 		} else {
// 			res = append(res, r)
// 		}
// 	}
//
// 	return res
// }
//
// func (f *CharacterForm) getAvailableGenders() []GenderShort {
// 	res := []GenderShort{}
//
// 	for _, g := range genders {
// 		if exclude, _ := util.InArray(f.Character.Race, g.DisallowedRaces); exclude {
// 			continue
// 		}
//
// 		if g.Only != "" && f.Character.Race != g.Only {
// 			continue
// 		}
//
// 		res = append(res, GenderShort{g.Name, g.ID})
// 	}
//
// 	return res
// }
//
// func (f *CharacterForm) getAvailableAestheticTraits() map[string]AestheticTraitCategoryShort {
// 	res := make(map[string]AestheticTraitCategoryShort)
//
// 	for _, cat := range aestheticTraitsCategorized { // Category Level
// 		if cat.Only != "" && f.Character.Race != cat.Only && f.Character.Gender != cat.Only {
// 			continue
// 		}
//
// 		if exclude, _ := util.InArray(f.Character.Race, cat.DisallowedRaces); exclude {
// 			continue
// 		}
//
// 		if exclude, _ := util.InArray(f.Character.Gender, cat.DisallowedGenders); exclude {
// 			continue
// 		}
//
// 		list := []AestheticTraitShort{}
//
// 		for _, t := range cat.Traits { // Trait inner loop
// 			if t.Only != "" && f.Character.Race != t.Only && f.Character.Gender != t.Only {
// 				continue
// 			}
//
// 			if exclude, _ := util.InArray(f.Character.Gender, t.DisallowedGenders); exclude {
// 				continue
// 			}
//
// 			if exclude, _ := util.InArray(f.Character.Race, t.DisallowedRaces); exclude {
// 				continue
// 			}
//
// 			list = append(list, t.Shorten())
// 		}
//
// 		res[cat.ID] = AestheticTraitCategoryShort{cat.Name, cat.Unique, cat.ID, list, cat.Minimum}
// 	}
//
// 	return res
// }
//
// func (f *CharacterForm) getAvailableFunctionalTraits() map[string]FunctionalTraitCategoryShort {
// 	res := make(map[string]FunctionalTraitCategoryShort)
//
// 	for _, cat := range functionalTraitsCategorized { // Category Level
// 		if cat.Only != "" && f.Character.Race != cat.Only && f.Character.Gender != cat.Only {
// 			continue
// 		}
//
// 		if exclude, _ := util.InArray(f.Character.Race, cat.DisallowedRaces); exclude {
// 			continue
// 		}
//
// 		if exclude, _ := util.InArray(f.Character.Gender, cat.DisallowedGenders); exclude {
// 			continue
// 		}
//
// 		list := []FunctionalTraitShort{}
//
// 		for _, t := range cat.Traits { // Trait inner loop
// 			if t.Only != "" && f.Character.Race != t.Only && f.Character.Gender != t.Only {
// 				continue
// 			}
//
// 			if exclude, _ := util.InArray(f.Character.Gender, t.DisallowedGenders); exclude {
// 				continue
// 			}
//
// 			if exclude, _ := util.InArray(f.Character.Race, t.DisallowedRaces); exclude {
// 				continue
// 			}
//
// 			list = append(list, t.Shorten())
// 		}
//
// 		res[cat.ID] = FunctionalTraitCategoryShort{cat.Name, cat.Unique, cat.ID, list, cat.Minimum}
// 	}
//
// 	return res
// }
//
// func (f *CharacterForm) getAvailableBackgrounds() []BackgroundShort {
// 	res := []BackgroundShort{}
//
// 	for _, bg := range backgrounds { // Category Level
// 		if bg.AllowAll {
// 			res = append(res, bg.Shorten())
// 			continue
// 		}
//
// 		// check against prerequisites to see if allowed @todo
// 	}
//
// 	return res
// }
