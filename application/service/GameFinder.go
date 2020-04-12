package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
)

type GameFinder struct {
	log    *log.Logger
	finder port_out.FindGame
}

func NewGameFinder(log *log.Logger, finder port_out.FindGame) *GameFinder {
	return &GameFinder{
		log: log,
		finder: finder,
	}
}

func (f *GameFinder) Find(code string) (*domain.Game, error) {
	g, err := f.finder.GetGameByCode(code)
	if err != nil {
		return nil, port_in.ErrGameNotFound
	}
	return g, nil
}
