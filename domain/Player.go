package domain

import (
	"time"

	"github.com/google/uuid"
)

type PlayerList []*Player

type PlayerID uuid.UUID

type Player struct {
	ID       PlayerID
	Username string
	Server   string
	JoinedAt time.Time
}

// NewPlayer creates new Player with username and server
func NewPlayer(username string, server string) *Player {
	return NewPlayerWithID(uuid.New(), username, server)
}

// NewPlayerWithID creates new Player with id, username and server
func NewPlayerWithID(id uuid.UUID, username string, server string) *Player {
	return &Player{ID: PlayerID(id), Username: username, Server: server}
}
