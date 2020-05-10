package query

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrGameNotFoundByCode = errors.New("game not found with code")
)

type FindGameByCode interface {
	FindGameByCode(code model.GameCode) (*model.Game, error)
}
