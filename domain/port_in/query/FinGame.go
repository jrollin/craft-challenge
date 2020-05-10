package query

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

type FindGame interface {
	FindGame(id model.GameID) (*model.Game, error)
}
