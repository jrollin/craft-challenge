package port_out

import "github.com/jrollin/craft-challenge/domain"

type StoreGame interface {
	StoreGame(game *domain.Game) error
}
