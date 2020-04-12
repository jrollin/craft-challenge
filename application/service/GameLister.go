package service

import (
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
)

type GameLister struct {
	log    *log.Logger
	lister port_out.ListGames
}

func NewGameLister(log *log.Logger, lister port_out.ListGames) *GameLister {
	return &GameLister{
		log:    log,
		lister: lister,
	}
}

func (f *GameLister) GetAllGames() ([]*domain.Game, error) {
	return f.lister.GetAllGames()
}
