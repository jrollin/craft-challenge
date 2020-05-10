package api

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"time"

	"github.com/google/uuid"
)

// swagger:response playerResponse
type Player struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Server   string    `json:"server"`
	JoinedAt time.Time `json:"joined_at"`
}

func NewPlayerFromDomain(player *model.Player) *Player {
	return &Player{
		ID:       uuid.UUID(player.ID),
		Username: player.Username,
		Server:   player.Server,
		JoinedAt: player.JoinedAt,
	}
}
