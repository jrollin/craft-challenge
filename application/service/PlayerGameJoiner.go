package service

import (
	"log"
	"time"

	"github.com/jrollin/craft-challenge/application/port_in/command"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type PlayerGameJoiner struct {
	l  *log.Logger
	ap port_out.AddPlayerToGame
	fg port_out.GetGameByCode
}

func NewPlayerGameJoiner(l *log.Logger, ap port_out.AddPlayerToGame, fg port_out.GetGameByCode) *PlayerGameJoiner {
	return &PlayerGameJoiner{l: l, ap: ap, fg: fg}
}

func (pg *PlayerGameJoiner) JoinPlayerToGame(joinGame *command.JoinGameCommand) error {
	// find existing game by code
	g, err := pg.fg.GetGameByCode(domain.GameCode(joinGame.Code))
	if err != nil {
		return port_out.ErrGameNotFound
	}

	// check game not started
	if g.IsStarted() {
		return command.ErrCannotJoinStartedGame
	}

	// check game not ended
	if g.IsEnded() {
		return command.ErrCannotJoinEndedGame
	}

	// @todo
	// check already exists username

	// check server not already exists

	p := &domain.Player{
		ID:       domain.PlayerID(joinGame.ID),
		Username: joinGame.Username,
		Server:   joinGame.Server,
		JoinedAt: time.Now(),
	}
	err = pg.ap.AddPlayerToGame(p, g)
	if err != nil {
		return command.ErrAddingPlayerToGameFailed
	}
	return nil

}
