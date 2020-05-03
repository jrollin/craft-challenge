package service

import (
	"log"

	"github.com/jrollin/craft-challenge/domain"
)

type GameStoryValidator struct {
	l *log.Logger
}

func NewGameStoryValidator(log *log.Logger) *GameStoryValidator {
	return &GameStoryValidator{
		l: log,
	}
}

func (v *GameStoryValidator) ValidateStory(story *domain.StoryValidationRequest) (*domain.StoryValidationResult, error) {
	return &domain.StoryValidationResult{}, nil
}
