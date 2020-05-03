package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPlayerNotFound             = errors.New("player not found")
	ErrUsernameAlreadyAssigned    = errors.New("username already assigned")
	ErrCannotDeleteNotfoundPlayer = errors.New("cannot delete unfound player")
	ErrGameAlreadyPublished       = errors.New("game already published")
	ErrGameAlreadyStarted         = errors.New("cannot start already started game")
	ErrGameAlreadyEnded           = errors.New("cannot end already ended game")
)

type GameList map[string]*Game

type GameCode string

type GameID uuid.UUID

type Game struct {
	ID          GameID
	Code        GameCode
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	StartedAt   time.Time
	EndedAt     time.Time
	Stories     Stories
	Players     map[PlayerID]*Player
}

// NewGame creates new Game with Code
func NewGame(code string) *Game {
	return NewGameWithID(uuid.New(), code)
}

// NewGameWithID creates a new Game with id and code
func NewGameWithID(id uuid.UUID, code string) *Game {
	return &Game{ID: GameID(id), Code: GameCode(code)}
}

func (g *Game) Start() error {
	if g.IsStarted() {
		return ErrGameAlreadyStarted
	}
	g.StartedAt = time.Now()

	return nil
}

func (g *Game) End() error {
	if g.IsEnded() {
		return ErrGameAlreadyEnded
	}
	g.EndedAt = time.Now()

	return nil
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

func (g *Game) PublishAt(date time.Time) error {
	if !g.PublishedAt.IsZero() {
		return ErrGameAlreadyPublished
	}
	g.PublishedAt = date
	return nil
}

func (g *Game) HaveEnoughPlayers() bool {
	if len(g.Players) > 0 {
		return true
	}
	return false
}

func (g *Game) LoadStories(stories Stories) error {
	g.Stories = stories
	return nil
}

func (g *Game) AddPlayer(player *Player) error {
	p, _ := g.GetPlayerByUsername(player.Username)

	if p != nil {
		return ErrUsernameAlreadyAssigned
	}

	g.Players[player.ID] = player

	return nil
}

func (g *Game) RemovePlayer(player *Player) error {
	p, err := g.GetPlayerByUsername(player.Username)
	if err != nil {
		return ErrCannotDeleteNotfoundPlayer
	}
	delete(g.Players, p.ID)

	return nil
}

func (g *Game) GetPlayerByUsername(username string) (*Player, error) {
	for _, p := range g.Players {
		if p.Username == username {
			return p, nil
		}
	}

	return nil, ErrPlayerNotFound
}
