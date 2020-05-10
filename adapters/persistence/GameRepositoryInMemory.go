package persistence

import (
	"errors"
	"fmt"
	"github.com/jrollin/craft-challenge/domain/model"
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain/port_out"
)

type GameRepositoryInMemory struct {
	GameList map[string]*model.Game
}

func NewGameRepositoryInMemoryAdapter() *GameRepositoryInMemory {
	gameList := make(map[string]*model.Game)
	gameList["abc"] = &model.Game{
		ID:        model.GameID(uuid.MustParse("11111111-2222-3333-4444-555555555555")),
		Code:      model.GameCode("abc"),
		CreatedAt: time.Now().UTC(),
		Players:   map[model.PlayerID]*model.Player{},
		Stories: []*model.Story{
			&model.Story{
				ID:          model.StoryID(uuid.New()),
				Title:       "First, a red bike",
				Description: "Build me a red bike",
				Specifications: []*model.Specification{
					&model.Specification{
						Description: "Bike should be red",
						Rules: []*model.Rule{
							&model.Rule{
								Description: "should have property color with value red",
								Query: model.Request{
									Method:       "GET",
									URL:          "/car/2",
									QueryParams:  nil,
									BodyParams:   nil,
									HeaderParams: nil,
								},
								Assertion: model.Assertion{
									RequestPart: "body",
									Matcher:     "equals",
									Param:       "$.color",
									Expected:    "red",
								},
							},
						},
					},
					&model.Specification{Description: "Bike should have 2 wheels"},
				},
			},
			&model.Story{
				ID:          model.StoryID(uuid.New()),
				Title:       "Then a yellow car",
				Description: "Build me a car !",
				Specifications: []*model.Specification{
					&model.Specification{Description: "Car should be yellow"},
					&model.Specification{Description: "Car can move forward"},
					&model.Specification{Description: "Car can move backward"},
				},
			},
		},
	}

	return &GameRepositoryInMemory{
		GameList: gameList,
	}
}

func (gr *GameRepositoryInMemory) GetAllGames() (model.GameList, error) {
	return gr.GameList, nil
}

func (gr *GameRepositoryInMemory) GetGameByCode(code model.GameCode) (*model.Game, error) {

	g, ok := gr.GameList[string(code)]
	if ok == false {
		return nil, port_out.ErrGameNotFound
	}
	return g, nil

}

func (gr *GameRepositoryInMemory) GetGame(id model.GameID) (*model.Game, error) {

	for _, p := range gr.GameList {
		if p.ID == id {
			return p, nil
		}
	}
	return nil, port_out.ErrGameNotFound

}

func (gr *GameRepositoryInMemory) AddGame(game *model.Game) error {
	_, ok := gr.GameList[string(game.Code)]
	if ok == true {
		return port_out.ErrGameAlreadyExists
	}
	gr.GameList[string(game.Code)] = game
	return nil
}

func (gr *GameRepositoryInMemory) AddPlayerToGame(player *model.Player, game *model.Game) error {
	g, err := gr.GetGameByCode(game.Code)
	if err != nil {
		return errors.New("error joining player to game")
	}

	fmt.Printf("game %s %s", g.Code, g.ID)

	gr.GameList[string(g.Code)].Players[player.ID] = player

	return nil
}

func (gr *GameRepositoryInMemory) ListGamePlayers(game *model.Game) (model.PlayerList, error) {
	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return nil, port_out.ErrGameNotFound
	}

	pl := []*model.Player{}

	for _, p := range g.Players {
		pl = append(pl, p)
	}

	return pl, nil
}

func (gr *GameRepositoryInMemory) StoreGame(game *model.Game) error {

	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return port_out.ErrGameNotFound
	}

	g.StartedAt = time.Now()

	return nil
}

func (gr *GameRepositoryInMemory) ListGameStories(game *model.Game) (model.Stories, error) {

	g, ok := gr.GameList[string(game.Code)]
	if ok == false {
		return nil, port_out.ErrGameNotFound
	}

	return g.Stories, nil
}
