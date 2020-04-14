package port_in

import (
	"github.com/jrollin/craft-challenge/domain"
)

type ListGamePlayers interface {
	ListGamePlayers(game *domain.Game) (domain.PlayerList, error)
}
