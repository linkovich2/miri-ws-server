package core

type Weather struct {
	Name  string
	Range PositionRange
	Type  WeatherType

	// needs some kind of timer, to help see when its going to stop, intesify or ebb
}

type WeatherType struct {
	Name        string
	Description string

	// @todo represent effects of weather on player character
}

func (w *Weather) getIntensity(p *Position) {
	// @stub, returns weather intensity based on a position
}

func (w *Weather) move(modifier string) {
	// @stub, move the "storm"
}

func (w *Weather) grow(modifier int) {
	// @stub, grow the "storm"
}

func (w *Weather) shrink(modifier int) {
	// @stub, shrink the "storm"
}

func (w *Weather) roomIn(r *Room) {
	// @stub check if this room is in this "storm"
}

func (w *Weather) resolveConflict(conflictingWeather *Weather) {
	// @stub, in case of overlapping weather
}

func (w *Weather) update() {
	// @stub, weather should decide things about itself
}

func generateWeather() {
	// @stub, generate weather from a starting point
}
