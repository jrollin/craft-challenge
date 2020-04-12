package domain

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	ID        uuid.UUID
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Stories   Stories
}

func NewGame(code string) *Game {
	return &Game{Code: code}
}

func (g *Game) LoadStories(stories Stories) error {
	g.Stories = stories
	return nil
}

func (g *Game) GetFirstStory() (*Story, error) {
	return g.Stories.GetFirstStory()
}
