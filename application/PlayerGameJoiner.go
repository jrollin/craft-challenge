package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"
	"time"

	"github.com/jrollin/craft-challenge/domain/port_in/command"
	"github.com/jrollin/craft-challenge/domain/port_out"
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
	g, err := pg.fg.GetGameByCode(model.GameCode(joinGame.Code))
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

	p := &model.Player{
		ID:       model.PlayerID(joinGame.ID),
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
