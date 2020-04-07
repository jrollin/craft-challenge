package port_in

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain"
)

var (
	ErrGameStorageInvalid = errors.New("Game is invalid")
	ErrGameStorageFailed = errors.New("Game storage failed")
)

type StoreGame interface {
	Store(game *domain.Game) error
}
