package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_in"
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
		return port_in.ErrGameMustHaveOneOrMorePlayers
	}

	if game.IsStarted() {
		return port_in.ErrGameAlreadyStarted
	}

	if game.IsEnded() {
		return port_in.ErrGameAlreadyEnded
	}

	return g.s.StartGame(game)
}
