package command

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/validator"
	"github.com/stretchr/testify/assert"
)

func TestAddGame_ValidationWithoutCode(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGameCommand(id, "")

	assert.Error(t, err)
	assert.NotEmpty(t, err.Error())

	assert.Len(t, err.(*validator.Error).Fields, 1)
}

func TestAddGame_Validation(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGameCommand(id, "abc")

	assert.NoError(t, err)
}
