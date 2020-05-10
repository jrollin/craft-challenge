package api

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"time"

	"github.com/google/uuid"
)

// swagger:response gameResponse
type Game struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

func NewGameFromDomain(game *model.Game) *Game {
	g := &Game{
		ID:          uuid.UUID(game.ID),
		Code:        string(game.Code),
		CreatedAt:   game.CreatedAt,
		UpdatedAt:   game.UpdatedAt,
		PublishedAt: game.PublishedAt,
	}

	return g
}
