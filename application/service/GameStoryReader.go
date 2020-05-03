package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
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

func (g *GameStoryReader) DisplayCurrentStoryForPlayer(game *domain.Game, player *domain.Player) (*domain.Story, error) {

	if !game.IsStarted() {
		return nil, port_in.ErrCannotDisplayStoryWhenGameNotStarted
	}

	if game.IsEnded() {
		return nil, port_in.ErrCannotDisplayStoryWhenGameEnded
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
