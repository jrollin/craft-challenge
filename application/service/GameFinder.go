package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameFinder struct {
	finder  port_out.FindGame
}

func NewGameFinder(finder port_out.FindGame) *GameFinder{
	return &GameFinder{
		finder,
	}
}

func (f *GameFinder) Find(code string) (*domain.Game, error) {
	if code == "111" {
		return nil, port_in.ErrGameNotFound
	}
	return f.finder.GetGameByCode(code)
}
