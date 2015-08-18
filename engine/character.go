package engine

import ()

/*Types that should be contained inside a character*/

type character struct {
}

type gender struct {
	Male   bool
	Female bool
}

type race struct {
	SpaghettiMonster bool
}

type aTraits struct {
}

type fTraits struct {
}

/*Functions, MapTo should correctly map each character to a specific type and return it in Json
depending on what they picked, ie SpaghettiMonster -> Meatballs and sauce,
or SpaghettiNoodle -> Floppy and wet*/

func (c *character) MapTo() {
}
