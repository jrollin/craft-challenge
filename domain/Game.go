package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrUsernameAlreadyAssigned    = errors.New("username already assigned")
	ErrCannotDeleteNotfoundPlayer = errors.New("cannot delete unfound player")
)

type GameList map[string]*Game

type Game struct {
	ID        uuid.UUID
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	StartedAt time.Time
	EndedAt   time.Time
	Stories   Stories
	Players   map[string]*Player
}

func NewGame(code string) *Game {
	return &Game{Code: code}
}

func (g *Game) IsStarted() bool {
	if g.StartedAt.IsZero() {
		return false
	}
	return true
}

func (g *Game) IsEnded() bool {
	if g.EndedAt.IsZero() {
		return false
	}

	now := time.Now().UTC()
	if g.EndedAt.Before(now) {
		return true
	}
	return false
}



func (g *Game) LoadStories(stories Stories) error {
	g.Stories = stories
	return nil
}

func (g *Game) GetFirstStory() (*Story, error) {
	return g.Stories.GetFirstStory()
}

func (g *Game) AddPlayer(player *Player) error {
	_, ok := g.Players[player.Username]
	if ok == true {
		return ErrUsernameAlreadyAssigned
	}

	g.Players[player.Username] = player

	return nil
}

func (g *Game) RemovePlayer(player *Player) error {
	_, ok := g.Players[player.Username]
	if ok == false {
		return ErrCannotDeleteNotfoundPlayer
	}
	delete(g.Players, player.Username)

	return nil
}
