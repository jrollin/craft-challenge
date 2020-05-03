package command

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var (
	ErrCannotJoinStartedGame    = errors.New("cannot join a started game")
	ErrCannotJoinEndedGame      = errors.New("cannot join an ended game")
	ErrAddingPlayerToGameFailed = errors.New("adding player to game failed")
)

type JoinPlayerToGame interface {
	JoinPlayerToGame(joinGame *JoinGameCommand) error
}

type JoinGameCommand struct {
	ID       uuid.UUID `validate:"required"`
	Code     string    `validate:"required"`
	Username string    `validate:"required"`
	Server   string    `validate:"required"`
}

func NewJoinGameCommand(ID uuid.UUID, code string, username string, server string) (*JoinGameCommand, error) {
	jg := &JoinGameCommand{ID: ID, Code: code, Username: username, Server: server}
	err := jg.validate()
	if err != nil {
		return nil, err
	}
	return jg, nil
}

func (jg *JoinGameCommand) validate() error {
	validate := validator.New()
	return validate.Struct(jg)
}
