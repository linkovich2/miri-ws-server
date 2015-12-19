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
	Disconnect  chan string
	Connections map[string]*Connection
}

func (game *Game) Start() {
	dice.SeedRandom() // seed rand for dice

	miri = &world.Miri
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, miri.Update) // start the world update loop

	for {
		select {
		case c := <-game.Input:
			logger.Write.Info("Connection [%s] sent Command: [%v]", c.Connection.ID, c.Value)
		case c := <-game.Connect:
			logger.Write.Info("Connection [%s] started in game with Character: [%s]", c.Socket.ID, c.Character.Name)
			game.Connections[c.Socket.ID] = c

			room := miri.Realms[c.Character.Realm].Rooms[c.Character.Position]
			msg := &response{
				Room:       room,
				Messages:   []string{"Connected"},
				Directions: GetAvailableDirections(&room, c.Character.Realm),
			}

			res, _ := json.Marshal(msg)
			c.Socket.Send(res)
		case c := <-game.Disconnect:
			// @todo perform other logging out actions
			delete(game.Connections, c)
		}
	}
}

func NewGame() *Game {
	return &Game{
		Input:       make(chan *Command),
		Connect:     make(chan *Connection),
		Disconnect:  make(chan string),
		Connections: make(map[string]*Connection),
	}
}
