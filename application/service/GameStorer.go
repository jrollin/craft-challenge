package service

import (
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
)

type GameStorer struct {
	storage port_out.StoreGame
}

func NewGameStorer(store port_out.StoreGame) *GameStorer {
	return &GameStorer{
		store,
	}
}

func (f *GameStorer) Store(game *domain.Game) error {
	// validate
	return f.storage.Store(game)
}
