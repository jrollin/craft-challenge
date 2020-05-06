package port_out

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFoundByCode = errors.New("Game not found by code")
)

type GetGameByCode interface {
	GetGameByCode(code domain.GameCode) (*domain.Game, error)
}
