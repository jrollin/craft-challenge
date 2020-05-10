package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"
	"time"

	"github.com/jrollin/craft-challenge/domain/port_in/command"
	"github.com/jrollin/craft-challenge/domain/port_out"
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

	game := &model.Game{
		ID:        model.GameID(AddGameCommand.ID),
		Code:      model.GameCode(AddGameCommand.Code),
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
