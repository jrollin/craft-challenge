package port_out

import (
	"github.com/jrollin/craft-challenge/domain/model"
)

type ListGamePlayers interface {
	ListGamePlayers(game *model.Game) (model.PlayerList, error)
}
