package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_in/command"
	"github.com/jrollin/craft-challenge/application/port_out"
)

type GameStarter struct {
	l *log.Logger
	g port_out.GetGame
	s port_out.StoreGame
}

func NewGameStarter(log *log.Logger, game port_out.GetGame, state port_out.StoreGame) *GameStarter {
	return &GameStarter{
		l: log,
		g: game,
		s: state,
	}
}

func (g *GameStarter) StartGame(cmd *command.StartGameCommand) error {

	game, err := g.g.GetGame(cmd.GameID)
	if err != nil {
		return port_out.ErrGameNotFound
	}

	if !game.HaveEnoughPlayers() {
		return command.ErrGameMustHaveOneOrMorePlayers
	}

	if game.IsStarted() {
		return command.ErrGameAlreadyStarted
	}

	if game.IsEnded() {
		return command.ErrGameAlreadyEnded
	}

	err = game.Start()
	if err != nil {
		return command.ErrGameStartFailed
	}

	return g.s.StoreGame(game)
}
