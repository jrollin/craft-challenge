package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGame_Start_WhenAlreadyStarted(t *testing.T) {
	g := &Game{StartedAt: time.Now()}
	err := g.Start()

	if err != ErrGameAlreadyStarted {
		t.Errorf("Cannot start an already started game %v", err)
	}
}

func TestGame_Start_WhenNotStarted(t *testing.T) {
	g := &Game{}
	err := g.Start()

	assert.Nil(t, err, "Can start a game")
}

func TestGame_End_WhenAlreadyEnded(t *testing.T) {
	g := &Game{EndedAt: time.Now()}
	err := g.End()

	if err != ErrGameAlreadyEnded {
		t.Errorf("Cannot end an already ended game %v", err)
	}
}

func TestGame_End_WhenNotEnded(t *testing.T) {
	g := &Game{}
	err := g.End()

	assert.Nil(t, err, "Can end a game")
}

func TestGame_IsStarted_WhenDateIsNotSet(t *testing.T) {
	g := &Game{}

	assert.False(t, g.IsStarted(), "Game should not be started if no value")
}

func TestGame_IsStarted_WhenDateIsSet(t *testing.T) {
	g := &Game{StartedAt: time.Now()}

	assert.True(t, g.IsStarted(), "Game should be started if no value")
}

func TestGame_IsEnded_WhenDateIsNotSet(t *testing.T) {
	g := &Game{}

	assert.False(t, g.IsEnded(), "Game should not be ended if no value")
}

func TestGame_IsEnded_WhenDateIsFuture(t *testing.T) {
	f := time.Now().AddDate(0, 0, 1)
	g := &Game{EndedAt: f}

	assert.False(t, g.IsEnded(), "Game should not be ended if in future")
}

func TestGame_IsEnded_WhenDateIsPast(t *testing.T) {
	f := time.Now().AddDate(0, 0, -1)
	g := &Game{EndedAt: f}

	assert.True(t, g.IsEnded(), "Game should be ended if in past")
}

func TestGame_PublishAt_WhenAlreadyExists(t *testing.T) {
	d := time.Now()
	g := &Game{PublishedAt: d}

	assert.Error(t, g.PublishAt(time.Now()), ErrGameAlreadyPublished)
}

func TestGame_PublishAt_WhenNotPublished(t *testing.T) {
	g := &Game{}
	d := time.Now()
	g.PublishAt(d)

	assert.EqualValues(t, d, g.PublishedAt)
}

func TestGame_GetPlayerByUsername_NotFound(t *testing.T) {

	players := make(map[PlayerID]*Player)
	p := &Game{
		Players: players,
	}

	_, err := p.GetPlayerByUsername("john")

	if err != ErrPlayerNotFound {
		t.Errorf("Game should return ErrPlayerNotFound %v", err)
	}
}

func TestGame_GetPlayerByUsername_Found(t *testing.T) {

	players := make(map[PlayerID]*Player)
	p1 := NewPlayerWithID(uuid.MustParse("a8ee8067-4732-4ca5-a188-7a76ef293e97"), "john", "")
	players[p1.ID] = p1
	p := &Game{
		Players: players,
	}

	found, err := p.GetPlayerByUsername("john")

	if found == nil {
		t.Errorf("Game should return player with username %v", err)
	}
}

func TestGame_AddPlayer_WhenUsernameAlreadyExists(t *testing.T) {

	players := make(map[PlayerID]*Player)
	p1 := NewPlayerWithID(uuid.MustParse("a8ee8067-4732-4ca5-a188-7a76ef293e97"), "john", "")
	players[p1.ID] = p1
	p := &Game{
		Players: players,
	}

	err := p.AddPlayer(NewPlayer("john", ""))

	if err != ErrUsernameAlreadyAssigned {
		t.Errorf("Cannot add player with same username %v", err)
	}
}

func TestGame_AddPlayer(t *testing.T) {

	players := make(map[PlayerID]*Player)
	p1 := NewPlayerWithID(uuid.MustParse("a8ee8067-4732-4ca5-a188-7a76ef293e97"), "john", "")
	players[p1.ID] = p1
	p := &Game{
		Players: players,
	}
	err := p.AddPlayer(NewPlayer("marc", ""))
	if err != nil {
		t.Errorf("Failed to add player : %v", err)
	}

	if len(p.Players) != 2 {
		t.Errorf("Players should be 2: %v", err)
	}
}

func TestGame_RemovePlayer_NotExists(t *testing.T) {
	players := make(map[PlayerID]*Player)
	p1 := NewPlayerWithID(uuid.MustParse("a8ee8067-4732-4ca5-a188-7a76ef293e97"), "john", "")
	players[p1.ID] = p1
	p := &Game{
		Players: players,
	}
	err := p.RemovePlayer(NewPlayer("marc", ""))
	if err != ErrCannotDeleteNotfoundPlayer {
		t.Errorf("Should not remove not found player : %v", err)
	}
}

func TestGame_RemovePlayer(t *testing.T) {
	players := make(map[PlayerID]*Player)
	p1 := NewPlayerWithID(uuid.MustParse("a8ee8067-4732-4ca5-a188-7a76ef293e97"), "john", "")
	players[p1.ID] = p1
	p := &Game{
		Players: players,
	}
	err := p.RemovePlayer(NewPlayer("john", ""))
	if err != nil {
		t.Errorf("Failed to remove player : %v", err)
	}
}
