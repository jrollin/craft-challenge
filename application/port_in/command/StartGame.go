package command

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameMustHaveOneOrMorePlayers = errors.New("game must have at least one player")
	ErrGameAlreadyStarted           = errors.New("game has already started")
	ErrGameAlreadyEnded             = errors.New("game has already ended")
)

type StartGame interface {
	StartGame(game *domain.Game) error
}
