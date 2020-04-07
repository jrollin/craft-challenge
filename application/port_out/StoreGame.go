package port_out

import "github.com/jrollin/craft-challenge/domain"

type StoreGame interface {
	Store(game *domain.Game) error
}
