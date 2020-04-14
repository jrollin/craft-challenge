package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
)

type GameFinder struct {
	l *log.Logger
	f port_out.FindGame
}

func NewGameFinder(log *log.Logger, finder port_out.FindGame) *GameFinder {
	return &GameFinder{
		l: log,
		f: finder,
	}
}

func (gf *GameFinder) FindByCode(code string) (*domain.Game, error) {
	g, err := gf.f.GetGameByCode(code)
	if err != nil {
		return nil, port_in.ErrGameNotFound
	}
	return g, nil
}
