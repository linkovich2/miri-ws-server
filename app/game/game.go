package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/util/dice"
	"github.com/jonathonharrell/miri-ws-server/app/util/filters"
	"time"
)

type Game struct {
	Input chan *Command
}

func (game *Game) Start() {
	dice.SeedRandom() // seed rand for dice

	miri := &core.World{"Miri", make(map[string]core.Realm)}             // load in the world, rooms, etc
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	for {
		select {
		case command := <-game.Input:
			logger.Write.Info("Received a command: %v", command.Value)
		}
	}
}

func NewGame() *Game {
	return &Game{make(chan *Command)}
}
