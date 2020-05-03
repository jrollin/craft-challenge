package port_out

import "github.com/jrollin/craft-challenge/domain"

type ListGameStories interface {
	ListGameStories(game *domain.Game) (domain.Stories, error)
}
