package api

import "github.com/jrollin/craft-challenge/domain"

type PlayerList []*Player

func NewPlayerListFromDomain(players domain.PlayerList) PlayerList {

	pl := PlayerList{}
	for _, p := range players {
		pl = append(pl, NewPlayerFromDomain(p))
	}

	return pl
}
