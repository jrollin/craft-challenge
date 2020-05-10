package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"

	"github.com/jrollin/craft-challenge/domain/port_in/query"
	"github.com/jrollin/craft-challenge/domain/port_out"
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

func (gf *GameFinder) FindGameByCode(code model.GameCode) (*model.Game, error) {
	g, err := gf.f.GetGameByCode(code)
	if err != nil {
		return nil, query.ErrGameNotFound
	}
	return g, nil
}
