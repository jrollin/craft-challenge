package port_out

import (
	"github.com/jrollin/craft-challenge/domain/model"
)

type ListGameStories interface {
	ListGameStories(game *model.Game) (model.Stories, error)
}
