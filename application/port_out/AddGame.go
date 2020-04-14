package port_out

import "github.com/jrollin/craft-challenge/domain"

type AddGame interface {
	AddGame(game *domain.Game) error
}
