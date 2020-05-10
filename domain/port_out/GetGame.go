package port_out

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrGameNotFound = errors.New("Game not found")
)

type GetGame interface {
	GetGame(GameID model.GameID) (*model.Game, error)
}
