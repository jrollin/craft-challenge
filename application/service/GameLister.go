package service

import (
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameLister struct {
	lister port_out.ListGames
}

func NewGameLister(lister port_out.ListGames) *GameLister {
	return &GameLister{
		lister,
	}
}

func (f *GameLister) GetAllGames() ([]*domain.Game, error) {
	return f.lister.GetAllGames()
}
