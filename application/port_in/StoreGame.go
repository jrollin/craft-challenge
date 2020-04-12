package port_in

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var (
	ErrGameStorageFailed = errors.New("game storage failed")
)

type StoreGame interface {
	Store(AddGame *AddGame) error
}

type AddGame struct {
	ID   uuid.UUID `validate:"required"`
	Code string    `validate:"required"`
}

func NewAddGame(ID uuid.UUID, code string) (*AddGame, error) {
	g := &AddGame{
		ID:   ID,
		Code: code}

	err := g.validate()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (a *AddGame) validate() error {
	validate := validator.New()
	return validate.Struct(a)
}



