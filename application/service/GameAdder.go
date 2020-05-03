package service

import (
	"log"
	"time"

	"github.com/jrollin/craft-challenge/application/port_in/command"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameAdder struct {
	l *log.Logger
	s port_out.AddGame
}

func NewGameAdder(log *log.Logger, store port_out.AddGame) *GameAdder {
	return &GameAdder{
		l: log,
		s: store,
	}
}

func (g *GameAdder) AddGame(AddGameCommand *command.AddGameCommand) error {

	game := &domain.Game{
		ID:        domain.GameID(AddGameCommand.ID),
		Code:      domain.GameCode(AddGameCommand.Code),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stories:   nil,
	}
	err := g.s.AddGame(game)
	if err != nil {
		g.l.Printf("[ERROR] Error while storing game %s", err)
		return command.ErrGameStorageFailed
	}
	return nil
}
