package port_out

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFound = errors.New("Game not found")
)

type FindGame interface {
	GetGameByCode(code domain.GameCode) (*domain.Game, error)
}
