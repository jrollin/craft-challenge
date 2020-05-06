package command

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain"
	"github.com/jrollin/craft-challenge/validator"
)

var (
	ErrGameMustHaveOneOrMorePlayers = errors.New("game must have at least one player")
	ErrGameAlreadyStarted           = errors.New("game has already started")
	ErrGameAlreadyEnded             = errors.New("game has already ended")
	ErrGameStartFailed              = errors.New("game start failed")
)

type StartGame interface {
	StartGame(cmd *StartGameCommand) error
}

type StartGameCommand struct {
	GameID domain.GameID `validate:"required"`
}

func NewStartGameCommand(id uuid.UUID) (*StartGameCommand, error) {
	c := &StartGameCommand{GameID: domain.GameID(id)}

	err := c.validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *StartGameCommand) validate() error {
	validate := validator.NewValidator()
	return validate.Validate(c)
}
