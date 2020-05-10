package port_out

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrGameAlreadyExists = errors.New("Game already exists")
)

type AddGame interface {
	AddGame(game *model.Game) error
}
