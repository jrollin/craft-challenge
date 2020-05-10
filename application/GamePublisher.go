package application

import (
	"log"

	"github.com/jrollin/craft-challenge/domain/port_in/command"
	"github.com/jrollin/craft-challenge/domain/port_out"
)

type GamePublisher struct {
	l *log.Logger
	g port_out.GetGame
	s port_out.StoreGame
}

func NewGamePublisher(log *log.Logger, game port_out.GetGame, state port_out.StoreGame) *GamePublisher {
	return &GamePublisher{
		l: log,
		g: game,
		s: state,
	}
}

func (g *GamePublisher) PublishGame(cmd *command.PublishGameCommand) error {

	game, err := g.g.GetGame(cmd.GameID)
	if err != nil {
		return port_out.ErrGameNotFound
	}

	game.PublishedAt = cmd.PublishedAt

	return g.s.StoreGame(game)
}
