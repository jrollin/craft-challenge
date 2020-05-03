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

func NewPlayer(username string, server string) *Player {
	return &Player{ID: PlayerID(uuid.New()), Username: username, Server: server}
}

func NewPlayerWithID(id uuid.UUID, username string, server string) *Player {
	p := NewPlayer(username, server)
	p.ID = PlayerID(id)
	return p
}
