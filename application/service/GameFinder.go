package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_in/query"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameFinder struct {
	l *log.Logger
	f port_out.GetGameByCode
}

func NewGameFinder(log *log.Logger, finder port_out.GetGameByCode) *GameFinder {
	return &GameFinder{
		l: log,
		f: finder,
	}
}

func (gf *GameFinder) FindGameByCode(code domain.GameCode) (*domain.Game, error) {
	g, err := gf.f.GetGameByCode(code)
	if err != nil {
		return nil, query.ErrGameNotFound
	}
	return g, nil
}
