package service

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/application/port_out"
	"github.com/jrollin/craft-challenge/domain"
	"log"
	"time"
)

type GameStorer struct {
	log *log.Logger
	store port_out.StoreGame
}

func NewGameStorer(log *log.Logger, store port_out.StoreGame) *GameStorer {
	return &GameStorer{
		log: log,
		store : store,
	}
}

func (f *GameStorer) Store(AddGame *port_in.AddGame) error {

	game := &domain.Game{
		ID:        AddGame.ID,
		Code:      AddGame.Code,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Stories:   nil,
	}
	err := f.store.Store(game)
	if err != nil {
		f.log.Printf("[ERROR] Error while storing game %s", err)
		return port_in.ErrGameStorageFailed
	}
	return nil
}
