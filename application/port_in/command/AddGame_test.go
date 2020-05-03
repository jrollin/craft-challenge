package command

import (
	"github.com/google/uuid"
	"testing"
)

func TestAddGameValidationWithoutCode(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGameCommand(id, "")

	if err == nil {
		t.Errorf("AddGameCommand Code is required")
	}
}

func TestAddGameValidation(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGameCommand(id, "abc")

	if err != nil {
		t.Errorf("AddGameCommand validation failed %v with %s", err, id)
	}
}
