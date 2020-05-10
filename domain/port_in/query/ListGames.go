package query

import (
	"github.com/jrollin/craft-challenge/domain/model"
)

type ListGames interface {
	GetAllGames() (model.GameList, error)
}
