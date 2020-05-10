package port_out

import (
	"github.com/jrollin/craft-challenge/domain/model"
)

type AddPlayerToGame interface {
	AddPlayerToGame(player *model.Player, game *model.Game) error
}
