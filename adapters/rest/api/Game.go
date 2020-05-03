package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain"
)

// swagger:response gameResponse
type Game struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGameFromDomain(game *domain.Game) *Game {
	g := &Game{
		ID:        uuid.UUID(game.ID),
		Code:      string(game.Code),
		CreatedAt: game.CreatedAt,
		UpdatedAt: game.UpdatedAt,
	}

	return g
}
