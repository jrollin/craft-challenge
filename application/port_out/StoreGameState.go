package port_out

import "github.com/jrollin/craft-challenge/domain"

type StoreGameState interface {
	StartGame(game *domain.Game) error
}
