package engine

import (
	"fmt"
	"time"

	"github.com/jonathonharrell/dice"
	"github.com/jonathonharrell/miri-ws-server/engine/websocket"
)

var TheWorld World

func Start() {
	dice.SeedRandom()
	go websocket.StartServer()

	// load in the world, rooms, etc
	TheWorld = NewWorld("The Miri")

	// start the world update loop
	go RunEvery(WORLD_UPDATE_LOOP_TIMER*time.Second, TheWorld.Update)

	var input string
	fmt.Scanln(&input)
}
