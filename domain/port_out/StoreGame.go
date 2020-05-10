package port_out

import (
	"github.com/jrollin/craft-challenge/domain/model"
)

type StoreGame interface {
	StoreGame(game *model.Game) error
}
