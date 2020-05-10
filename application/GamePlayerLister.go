package application

import (
	"github.com/jrollin/craft-challenge/domain/model"
	"log"

	"github.com/jrollin/craft-challenge/domain/port_out"
)

type GamePlayerLister struct {
	l   *log.Logger
	lgp port_out.ListGamePlayers
}

func NewGamePlayerLister(log *log.Logger, lgp port_out.ListGamePlayers) *GamePlayerLister {
	return &GamePlayerLister{
		l:   log,
		lgp: lgp,
	}
}

func (g *GamePlayerLister) ListGamePlayers(game *model.Game) (model.PlayerList, error) {
	return g.lgp.ListGamePlayers(game)
}
