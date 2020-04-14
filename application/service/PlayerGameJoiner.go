package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
	"time"
)

type PlayerGameJoiner struct {
	l  *log.Logger
	ap port_out.AddPlayerToGame
	fg port_out.FindGame
}

func NewPlayerGameJoiner(l *log.Logger, ap port_out.AddPlayerToGame, fg port_out.FindGame) *PlayerGameJoiner {
	return &PlayerGameJoiner{l: l, ap: ap, fg: fg}
}

func (pg *PlayerGameJoiner) JoinPlayerToGame(joinGame *port_in.JoinGameCommand) error {
	// find existing game by code
	g, err := pg.fg.GetGameByCode(joinGame.Code)
	if err != nil {
		return port_in.ErrGameNotFound
	}

	// check game not started
	if g.IsStarted() {
		return port_in.ErrCannotJoinStartedGame
	}

	// check game not ended
	if g.IsEnded() {
		return port_in.ErrCannotJoinEndedGame
	}

	// @todo
	// check already exists username

	// check server not already exists

	p := &domain.Player{
		ID:       joinGame.ID,
		Username: joinGame.Username,
		Server:   joinGame.Server,
		JoinedAt: time.Now(),
	}
	err = pg.ap.AddPlayerToGame(p, g)
	if err != nil {
		return port_in.ErrAddingPlayerToGameFailed
	}
	return nil

}
