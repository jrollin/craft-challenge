package service

import (
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
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

func (g *GamePlayerLister) ListGamePlayers(game *domain.Game) (domain.PlayerList, error) {
	return g.lgp.ListGamePlayers(game)
}
