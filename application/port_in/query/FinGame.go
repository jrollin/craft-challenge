package query

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

type FindGame interface {
	FindGame(id domain.GameID) (*domain.Game, error)
}
