package api

import (
	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain"
	"time"
)

type Game struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGame(game *domain.Game) *Game {

	g := &Game{
		ID:        game.ID,
		Code:      game.Code,
		CreatedAt: game.CreatedAt,
		UpdatedAt: game.UpdatedAt,
	}

	return g
}
