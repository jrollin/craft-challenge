package command

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/validator"
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
	validate := validator.NewValidator()
	return validate.Validate(a)
}
