package command

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var (
	ErrGameStorageFailed = errors.New("game storage failed")
)

type AddGame interface {
	AddGame(AddGameCommand *AddGameCommand) error
}

type AddGameCommand struct {
	ID   uuid.UUID `validate:"required"`
	Code string    `validate:"required"`
}

func NewAddGameCommand(ID uuid.UUID, code string) (*AddGameCommand, error) {
	g := &AddGameCommand{
		ID:   ID,
		Code: code}

	err := g.validate()
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (a *AddGameCommand) validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
