package query

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFoundByCode = errors.New("game not found with code")
)

type FindGameByCode interface {
	FindGameByCode(code domain.GameCode) (*domain.Game, error)
}
