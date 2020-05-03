package port_in

import "github.com/jrollin/craft-challenge/domain"

type ValidateStory interface {
	ValidateStory(story *domain.StoryValidationRequest) (*domain.StoryValidationResult, error)
}
