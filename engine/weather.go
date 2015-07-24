package engine

import ()

type Weather struct {
  Name string
  Range PositionRange
  Type WeatherType

  // needs some kind of timer, to help see when its going to stop, intesify or ebb
}

type WeatherType struct {
  Name string
  Description string

  // @todo represent effects of weather on player character
}

func (weather *Weather) GetIntensity(pos *Position) {
  // @stub, returns weather intensity based on a position
}

func (weather *Weather) Move(modifier string) {
  // @stub, move the "storm"
}

func (weather *Weather) Grow(modifier int) {
  // @stub, grow the "storm"
}

func (weather *Weather) Shrink(modifier int) {
  // @stub, shrink the "storm"
}

func (weather *Weather) RoomIn(room *Room) {
  // @stub check if this room is in this "storm"
}

func (weather *Weather) ResolveConflict(conflictingWeather *Weather) {
  // @stub, in case of overlapping weather
}

func (weather *Weather) Update() {
  // @stub, weather should decide things about itself
}

func GenerateWeather() {
  // @stub, generate weather from a starting point
}
