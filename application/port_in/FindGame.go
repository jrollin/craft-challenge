package port_in

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFound = errors.New("game not found")
)

type FindGame interface {
	FindByCode(code domain.GameCode) (*domain.Game, error)
}
