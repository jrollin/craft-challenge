package port_out

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrGameNotFoundByCode = errors.New("Game not found by code")
)

type GetGameByCode interface {
	GetGameByCode(code model.GameCode) (*model.Game, error)
}
