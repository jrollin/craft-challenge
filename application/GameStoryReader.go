package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"

	"github.com/jrollin/craft-challenge/domain/port_in/query"
	"github.com/jrollin/craft-challenge/domain/port_out"
)

type GameStoryReader struct {
	l *log.Logger
	s port_out.ListGameStories
}

func NewGameStoryReader(log *log.Logger, stories port_out.ListGameStories) *GameStoryReader {
	return &GameStoryReader{
		l: log,
		s: stories,
	}
}

func (g *GameStoryReader) DisplayCurrentStoryForPlayer(game *model.Game, player *model.Player) (*model.Story, error) {

	if !game.IsStarted() {
		return nil, query.ErrCannotDisplayStoryWhenGameNotStarted
	}

	if game.IsEnded() {
		return nil, query.ErrCannotDisplayStoryWhenGameEnded
	}

	// retrieves stories for game
	stories, err := g.s.ListGameStories(game)
	if err != nil {
		return nil, err
	}

	// @todo
	// get last validated story for player and display next

	story := stories[0]

	return story, nil
}
