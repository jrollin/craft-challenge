package command

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewStartGameCommand_validate(t *testing.T) {
	id := uuid.New()
	_, err := NewStartGameCommand(id)

	assert.NoError(t, err)
}
