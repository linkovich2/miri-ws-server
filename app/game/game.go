package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/content/world"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/util/dice"
	// "github.com/jonathonharrell/miri-ws-server/app/util/filters"
	"encoding/json"
	"time"
)

var miri *core.World

type Game struct {
	Input       chan *Command
	Connect     chan *Connection
	Connections map[string]*Connection
}

func (game *Game) Start() {
	dice.SeedRandom() // seed rand for dice

	miri = &world.Miri
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	for {
		select {
		case c := <-game.Input:
			logger.Write.Info("Received a command: %v", c.Value)
		case c := <-game.Connect:
			logger.Write.Info("Connection [%s] started in game with Character: [%s]", c.Socket.ID, c.Character.Name)
			game.Connections[c.Socket.ID] = c
			logger.Write.Info("Num connections: %v", len(game.Connections))

			room := miri.Realms[c.Character.Realm].Rooms[c.Character.Position]
			msg := &response{
				Room:       room,
				Messages:   []string{"Connected"},
				Directions: GetAvailableDirections(&room, c.Character.Realm),
			}

			res, _ := json.Marshal(msg)
			c.Socket.Send(res)
		}
	}
}

func NewGame() *Game {
	return &Game{make(chan *Command), make(chan *Connection), make(map[string]*Connection)}
}
