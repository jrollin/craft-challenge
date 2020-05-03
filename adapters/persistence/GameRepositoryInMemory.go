package persistence

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
)

type GameRepositoryInMemory struct {
	GameList map[string]*domain.Game
}

func NewGameRepositoryInMemoryAdapter() *GameRepositoryInMemory {
	gameList := make(map[string]*domain.Game)
	gameList["abc"] = &domain.Game{
		ID:        uuid.New(),
		Code:      domain.GameCode("abc"),
		CreatedAt: time.Now().UTC(),
		Players:   map[domain.PlayerID]*domain.Player{},
		Stories: []*domain.Story{
			&domain.Story{
				ID:          domain.StoryID(uuid.New()),
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
				ID:          domain.StoryID(uuid.New()),
				Title:       "Then a yellow car",
				Description: "Build me a car !",
				Specifications: []*domain.Specification{
					&domain.Specification{Description: "Car should be yellow"},
					&domain.Specification{Description: "Car can move forward"},
					&domain.Specification{Description: "Car can move backward"},
				},
			},
		},
	}

	return &GameRepositoryInMemory{
		GameList: gameList,
	}
}

func (gr *GameRepositoryInMemory) GetAllGames() (domain.GameList, error) {
	return gr.GameList, nil
}

func (gr *GameRepositoryInMemory) GetGameByCode(code domain.GameCode) (*domain.Game, error) {

	g, ok := gr.GameList[string(code)]
	if ok == false {
		return nil, port_in.ErrGameNotFound
	}
	return g, nil

}

func (gr *GameRepositoryInMemory) AddGame(game *domain.Game) error {
	_, ok := gr.GameList[string(game.Code)]
	if ok == true {
		return port_in.ErrGameStorageFailed
	}
	gr.GameList[string(game.Code)] = game
	return nil
}

func (gr *GameRepositoryInMemory) AddPlayerToGame(player *domain.Player, game *domain.Game) error {
	g, err := gr.GetGameByCode(game.Code)
	if err != nil {
		return errors.New("error joining player to game")
	}

	fmt.Printf("game %s %s", g.Code, g.ID)

	gr.GameList[string(g.Code)].Players[player.ID] = player

	return nil
}

func (gr *GameRepositoryInMemory) ListGamePlayers(game *domain.Game) (domain.PlayerList, error) {
	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return nil, port_in.ErrGameNotFound
	}

	pl := []*domain.Player{}

	for _, p := range g.Players {
		pl = append(pl, p)
	}

	return pl, nil
}

func (gr *GameRepositoryInMemory) StartGame(game *domain.Game) error {

	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return port_in.ErrGameNotFound
	}

	g.StartedAt = time.Now()

	return nil
}

func (gr *GameRepositoryInMemory) ListGameStories(game *domain.Game) (domain.Stories, error) {

	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return nil, port_in.ErrGameNotFound
	}

	return g.Stories, nil
}
