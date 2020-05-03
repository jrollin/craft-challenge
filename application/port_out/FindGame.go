package port_out

import "github.com/jrollin/craft-challenge/domain"

type FindGame interface {
	GetGameByCode(code domain.GameCode) (*domain.Game, error)
}
