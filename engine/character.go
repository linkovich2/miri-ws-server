package engine

import ()

/*character should be something every connection has?*/
type character struct {
	genders map[*gender]bool
	races   map[*race]bool
	aTraits map[*aTrait]bool
	fTraits map[*fTrait]bool
}

/*Types that should be contained inside a character*/
type gender struct {
	Male   bool
	Female bool
}

type race struct {
	SpaghettiMonster bool
}

type aTrait struct {
}

type fTrait struct {
}

/*Methods, MapTo should correctly map each character to specific types,
depending on what they picked. ie SpaghettiMonster -> Meatballs, or sauce,
or SpaghettiNoodle -> Floppy, or wet*/

func (c *character) MapTo() {
}
