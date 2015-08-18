package engine

import ()

/*Types that should be contained inside a character*/

type character struct {
	genders map[*gender]bool
	races   map[*race]bool
	aTraits map[*aTrait]bool
	fTraits map[*fTrait]bool
}

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

/*Functions, MapTo should correctly map each character to specific types,
depending on what they picked. ie SpaghettiMonster -> Meatballs, or sauce,
or SpaghettiNoodle -> Floppy, or wet*/

func (c *character) MapTo() {
}
