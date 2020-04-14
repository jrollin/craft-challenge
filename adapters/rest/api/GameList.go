package api

import (
	"github.com/jrollin/craft-challenge/domain"
)

type GameList struct {
	Games []*Game `json:"games"`
}

func NewGameListFromDomain(games domain.GameList) *GameList {
	gl := &GameList{}
	for _, g := range games {
		gl.Games = append(gl.Games, NewGameFromDomain(g))
	}

	return gl
}
