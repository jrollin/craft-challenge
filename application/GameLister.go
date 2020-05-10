package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"

	"github.com/jrollin/craft-challenge/domain/port_out"
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

func (gl *GameLister) GetAllGames() (model.GameList, error) {
	return gl.lg.GetAllGames()
}
