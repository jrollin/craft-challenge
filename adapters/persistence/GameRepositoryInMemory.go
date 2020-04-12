package persistence

import (
	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
	"time"
)


var gameList = []*domain.Game{
	&domain.Game{
		ID:        uuid.New(),
		Code:      "abc",
		CreatedAt: time.Now().UTC(),
		Stories: []*domain.Story{
			&domain.Story{
				ID:          "1",
				Title:       "First, a red bike",
				Description: "Build me a red bike",
				Specifications: []*domain.Specification{
					&domain.Specification{
						Description: "Bike should be red",
						Rules: []*domain.Rule{
							&domain.Rule{
								Description: "should have property color with value red",
								Query: domain.Request{
									Method:       "GET",
									URL:          "/car/2",
									QueryParams:  nil,
									BodyParams:   nil,
									HeaderParams: nil,
								},
								Assertion: domain.Assertion{
									RequestPart: "body",
									Matcher:     "equals",
									Param:       "$.color",
									Expected:    "red",
								},
							},
						},
					},
					&domain.Specification{Description: "Bike should have 2 wheels"},
				},
			},
			&domain.Story{
				ID:          "2",
				Title:       "Then a yellow car",
				Description: "Build me a car !",
				Specifications: []*domain.Specification{
					&domain.Specification{Description: "Car should be yellow"},
					&domain.Specification{Description: "Car can move forward"},
					&domain.Specification{Description: "Car can move backward"},
				},
			},
		},
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



