package port_out

import "github.com/jrollin/craft-challenge/domain"

type FindGame interface {
	GetGameByCode(code string) (*domain.Game, error)
}
