package port_in

import (
	"github.com/google/uuid"
	"testing"
)

func TestAddGameValidationWithoutCode(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGame(id, "")

	if err == nil {
		t.Errorf("AddGame Code is required")
	}
}

func TestAddGameValidation(t *testing.T) {
	id := uuid.New()
	_, err := NewAddGame(id, "abc")

	if err != nil {
		t.Errorf("AddGame validation failed %v with %s", err, id)
	}
}
