package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
	"time"
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

func (g *GameAdder) AddGame(AddGameCommand *port_in.AddGameCommand) error {

	game := &domain.Game{
		ID:        AddGameCommand.ID,
		Code:      AddGameCommand.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stories:   nil,
	}
	err := g.s.AddGame(game)
	if err != nil {
		g.l.Printf("[ERROR] Error while storing game %s", err)
		return port_in.ErrGameStorageFailed
	}
	return nil
}