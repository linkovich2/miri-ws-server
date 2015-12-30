package game

import (
	"github.com/jonathonharrell/miri-ws-server/app/content/world"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	db "github.com/jonathonharrell/miri-ws-server/app/persistence"
	"github.com/jonathonharrell/miri-ws-server/app/server"
	"github.com/jonathonharrell/miri-ws-server/app/util"
	"github.com/jonathonharrell/miri-ws-server/app/util/dice"
	// "github.com/jonathonharrell/miri-ws-server/app/util/filters"
	"encoding/json"
	"time"
)

type Game struct {
	Input       chan *Command
	Connect     chan *Connection
	Disconnect  chan string
	Connections map[string]*Connection
	World       *core.World
}

func (game *Game) Start() {
	dice.SeedRandom() // seed rand for dice

	game.World = &world.Miri
	game.World.SetSendCallback(func(id, message string) {
		if connection, exists := game.Connections[id]; exists {
			res, _ := json.Marshal(&response{Messages: []string{message}})
			connection.Socket.Send(res)
		}
	})
	go util.RunEvery(core.WorldUpdateLoopTimer*time.Second, game.World.Update) // start the world update loop

	for {
		select {
		case c := <-game.Input:
			game.handleInput(c)
		case c := <-game.Connect:
			game.handleConnection(c)
		case c := <-game.Disconnect:
			game.handleDisconnection(c)
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

func (game *Game) handleConnection(c *Connection) {
	logger.Write.Info("Connection [%s] started in game with Character: [%s]", c.Socket.ID, c.Character.Name)
	game.Connections[c.Socket.ID] = c
	room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
	room.Add(c.Socket.ID)
	game.World.Realms[c.Character.Realm].Rooms[c.Character.Position] = room
	game.defaultMessage(c.Socket, c.Character, []string{"<system>Connected;</system>"})
}

func (game *Game) handleDisconnection(conn string) {
	if c, exists := game.Connections[conn]; exists {
		go func() {
			db.SaveCharacter(c.Character)
			err := c.Character.AddState(core.StateLoggingOut)
			if err != nil {
				logger.Write.Error(err.Error()) // @todo add some detail here
			}

			time.Sleep(10 * time.Second) // wait for 10 seconds then log out the character

			c.Character.RemoveState(core.StateLoggingOut)
			room := game.World.Realms[c.Character.Realm].Rooms[c.Character.Position]
			room.Remove(c.Socket.ID)
			delete(game.Connections, c.Socket.ID)
		}()
	}
}

func (game *Game) handleInput(c *Command) {
	if f, exists := commandRegistrar[c.Value]; exists {
		f(game, c)
	} else {
		logger.Write.Error("Connection [%s] sent Command: [%v], but it doesn't exist.", c.Connection.ID, c.Value)
	}
}

func (game *Game) simpleMessage(s *server.Connection, messages []string) {
	msg := miniResponse{Messages: messages}
	res, _ := json.Marshal(&msg)
	s.Send(res)
}

func (game *Game) defaultMessage(s *server.Connection, c *core.Character, messages []string) {
	msg := response{Messages: messages, State: c.GetStateString()}
	if value, exists := game.World.Realms[c.Realm].Rooms[c.Position]; exists {
		msg.Room = value
		msg.Directions = game.getAvailableDirections(value, c.Realm)
	}

	res, _ := json.Marshal(&msg)
	s.Send(res)
}

func (game *Game) broadcastToRoom(originator *server.Connection, character *core.Character, message string, messageForOriginator string, room *core.Room) {
	game.defaultMessage(originator, character, []string{messageForOriginator})

	for _, id := range room.Connections {
		if id != originator.ID {
			c := game.Connections[id]
			game.simpleMessage(c.Socket, []string{message})
		}
	}
}

func (game *Game) CurrentlyPlaying(c *core.Character) bool {
	for _, connection := range game.Connections {
		if connection.Character.ID == c.ID {
			return true
		}
	}

	return false
}
