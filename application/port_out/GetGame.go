package port_out

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFound = errors.New("Game not found")
)

type GetGame interface {
	GetGame(GameID domain.GameID) (*domain.Game, error)
}
