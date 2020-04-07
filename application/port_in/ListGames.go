package port_in

import (
	"github.com/jrollin/craft-challenge/domain"
)

type ListGames interface {
	GetAllGames() ([]*domain.Game, error)
}
