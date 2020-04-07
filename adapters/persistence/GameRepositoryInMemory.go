package persistence

import (
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
	"time"
)



var gameList = []*domain.Game{
	&domain.Game{
		Id: 1,
		Code:      "abc",
		CreatedAt: time.Now().UTC(),
	},
	&domain.Game{
		Id: 2,
		Code:      "def",
		CreatedAt: time.Now().UTC(),
	},
}


type GameRepositoryInMemory struct {
	
}

func NewGameRepositoryInMemoryAdapter() *GameRepositoryInMemory {
	return &GameRepositoryInMemory{}
}

func (g *GameRepositoryInMemory) GetAllGames() ([]*domain.Game, error) {
	return gameList, nil
}


func (g *GameRepositoryInMemory) GetGameByCode(code string) (*domain.Game, error) {
	for _, p := range gameList {
		if p.Code == code {
			return p, nil
		}
	}
	return nil, port_in.ErrGameNotFound

}

func (g *GameRepositoryInMemory) Store(game *domain.Game) error {
	for _, p := range gameList {
		if p.Code == game.Code {
			return port_in.ErrGameStorageFailed
		}
	}
	gameList = append(gameList, game)
	return nil
}



