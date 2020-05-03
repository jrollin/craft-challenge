package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain"
)

// swagger:response playerResponse
type Player struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Server   string    `json:"server"`
	JoinedAt time.Time `json:"joined_at"`
}

func NewPlayerFromDomain(player *domain.Player) *Player {
	return &Player{
		ID:       uuid.UUID(player.ID),
		Username: player.Username,
		Server:   player.Server,
		JoinedAt: player.JoinedAt,
	}
}
