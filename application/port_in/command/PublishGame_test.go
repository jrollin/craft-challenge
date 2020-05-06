package command

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/validator"
	"github.com/stretchr/testify/assert"
)

func TestPublishGameCommand_validate(t *testing.T) {
	id := uuid.New()
	_, err := NewPublishGameCommand(id, time.Now())

	assert.NoError(t, err)
}

func TestPublishGameCommand_ValidationFails(t *testing.T) {
	_, err := NewPublishGameCommand(uuid.UUID{}, time.Time{})

	assert.Error(t, err)
	assert.NotEmpty(t, err.Error())

	assert.Len(t, err.(*validator.Error).Fields, 2)
}
