package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/util/dice"
	// "github.com/jonathonharrell/miri-ws-server/app/util/filters"
	"time"
)

type Game struct {
	Input       chan *Command
	Connect     chan *Connection
	Connections map[string]*Connection
}

func (game *Game) Start() {
	dice.SeedRandom() // seed rand for dice

	miri := &core.World{"Miri", make(map[string]core.Realm)}             // load in the world, rooms, etc
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	for {
		select {
		case c := <-game.Input:
			logger.Write.Info("Received a command: %v", c.Value)
		case c := <-game.Connect:
			logger.Write.Info("Connection [%s] started in game with Character: [%s]", c.Socket.ID, c.Character.Name)
			game.Connections[c.Socket.ID] = c

			// @todo we should send the initial play info after this
		}
	}
}

func NewGame() *Game {
	return &Game{make(chan *Command), make(chan *Connection), make(map[string]*Connection)}
}
