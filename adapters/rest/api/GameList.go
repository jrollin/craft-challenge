package api

import (
	"github.com/jrollin/craft-challenge/domain"
)

type GameList struct {
	Games []*Game `json:"games"`
}

func NewGameList(games []*domain.Game) *GameList {

	gl := &GameList{}

	for _, p := range games {
		gl.Games = append(gl.Games, NewGame(p))
	}

	return gl
}
