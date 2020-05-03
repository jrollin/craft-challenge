package command

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/domain"
	"github.com/jrollin/craft-challenge/validator"
)

var (
	ErrGamePublicationError = errors.New("game publication error")
)

// PublishGame defines how to publish game
type PublishGame interface {
	publishGame(cmd *PublishGameCommand) error
}

// PublishGameCommand defines command to use
type PublishGameCommand struct {
	GameID      domain.GameID `validate:"required"`
	PublishedAt time.Time     `validate:"required"`
}

// NewPublishGameCommand creates a new PublishGameCommand
func NewPublishGameCommand(GameID uuid.UUID, date time.Time) (*PublishGameCommand, error) {
	c := &PublishGameCommand{GameID: domain.GameID(GameID), PublishedAt: date}

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
