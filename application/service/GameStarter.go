package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_in/command"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameStarter struct {
	l *log.Logger
	s port_out.StoreGameState
}

func NewGameStarter(log *log.Logger, state port_out.StoreGameState) *GameStarter {
	return &GameStarter{
		l: log,
		s: state,
	}
}

func (g *GameStarter) StartGame(game *domain.Game) error {

	if !game.HaveEnoughPlayers() {
		return command.ErrGameMustHaveOneOrMorePlayers
	}

	if game.IsStarted() {
		return command.ErrGameAlreadyStarted
	}

	if game.IsEnded() {
		return command.ErrGameAlreadyEnded
	}

	return g.s.StartGame(game)
}
