package port_in

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameNotFound = errors.New("Game not found")
)

type FindGame interface {
	Find(code string) (*domain.Game, error)
}
