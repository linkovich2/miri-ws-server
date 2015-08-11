package engine

type weather struct {
	name  string
	rng positionRange
	tp  weatherType

	// needs some kind of timer, to help see when its going to stop, intesify or ebb
}

type weatherType struct {
	name        string
	description string

	// @todo represent effects of weather on player character
}

func (w *weather) getIntensity(p *position) {
	// @stub, returns weather intensity based on a position
}

func (w *weather) move(modifier string) {
	// @stub, move the "storm"
}

func (w *weather) grow(modifier int) {
	// @stub, grow the "storm"
}

func (w *weather) shrink(modifier int) {
	// @stub, shrink the "storm"
}

func (w *weather) roomIn(r *room) {
	// @stub check if this room is in this "storm"
}

func (w *weather) resolveConflict(conflictingWeather *weather) {
	// @stub, in case of overlapping weather
}

func (w *weather) update() {
	// @stub, weather should decide things about itself
}

func generateWeather() {
	// @stub, generate weather from a starting point
}
