package query

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
)

var (
	ErrCannotDisplayStoryWhenGameNotStarted = errors.New("game is not started")
	ErrCannotDisplayStoryWhenGameEnded      = errors.New("game has ended")
)

type DisplayCurrentStory interface {
	DisplayCurrentStoryForPlayer(game *model.Game, player *model.Player) (*model.Story, error)
}
