package service

import (
	"log"

	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameLister struct {
	l  *log.Logger
	lg port_out.ListGames
}

func NewGameLister(log *log.Logger, lg port_out.ListGames) *GameLister {
	return &GameLister{
		l:  log,
		lg: lg,
	}
}

func (gl *GameLister) GetAllGames() (domain.GameList, error) {
	return gl.lg.GetAllGames()
}
