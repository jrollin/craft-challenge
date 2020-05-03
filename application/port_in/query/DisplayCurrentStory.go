package query

import (
	"errors"

	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrCannotDisplayStoryWhenGameNotStarted = errors.New("game is not started")
	ErrCannotDisplayStoryWhenGameEnded      = errors.New("game has ended")
)

type DisplayCurrentStory interface {
	DisplayCurrentStoryForPlayer(game *domain.Game, player *domain.Player) (*domain.Story, error)
}
