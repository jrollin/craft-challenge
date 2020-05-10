package command

import (
	"errors"
	"github.com/jrollin/craft-challenge/domain/model"
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/validator"
)

var (
	ErrGamePublicationError = errors.New("game publication error")
)

// PublishGame defines how to publish game
type PublishGame interface {
	PublishGame(cmd *PublishGameCommand) error
}

// PublishGameCommand defines command to use
type PublishGameCommand struct {
	GameID      model.GameID `validate:"required"`
	PublishedAt time.Time    `validate:"required"`
}

// NewPublishGameCommand creates a new PublishGameCommand
func NewPublishGameCommand(GameID uuid.UUID, date time.Time) (*PublishGameCommand, error) {
	c := &PublishGameCommand{GameID: model.GameID(GameID), PublishedAt: date}

	err := c.validate()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (pg *PublishGameCommand) validate() error {
	validate := validator.NewValidator()
	return validate.Validate(pg)
}
