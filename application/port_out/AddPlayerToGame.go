package port_out

import "github.com/jrollin/craft-challenge/domain"

type AddPlayerToGame interface {
	AddPlayerToGame(player *domain.Player, game *domain.Game) error
}
