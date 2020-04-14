package domain

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	ID       uuid.UUID
	Username string
	Server   string
	JoinedAt time.Time
}


type PlayerList []*Player
