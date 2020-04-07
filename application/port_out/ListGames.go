package port_out

import "github.com/jrollin/craft-challenge/domain"

type ListGames interface {
	GetAllGames() ([]*domain.Game, error)
}
