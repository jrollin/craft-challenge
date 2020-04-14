package domain

import "testing"

func TestGameAddPlayerWhenUsernameAlreadyExists(t *testing.T) {

	players := make(map[string]*Player)
	players["john"] = &Player{Username: "john"}
	p := &Game{
		Players: players,
	}

	err := p.AddPlayer(&Player{Username: "john"})

	if err != ErrUsernameAlreadyAssigned {
		t.Errorf("Cannot add player with same username %v", err)
	}
}

func TestGameAddPlayer(t *testing.T) {

	players := make(map[string]*Player)
	players["john"] = &Player{Username: "john"}
	p := &Game{
		Players: players,
	}
	err := p.AddPlayer(&Player{Username: "marc"})
	if err != nil {
		t.Errorf("Failed to add player : %v", err)
	}

	if len(p.Players) != 2 {
		t.Errorf("Players should be 2: %v", err)
	}
}

func TestGameRemovePlayerNotExists(t *testing.T) {
	players := make(map[string]*Player)
	players["john"] = &Player{Username: "john"}
	p := &Game{
		Players: players,
	}
	err := p.RemovePlayer(&Player{Username: "marc"})
	if err != ErrCannotDeleteNotfoundPlayer {
		t.Errorf("Failed to remove player : %v", err)
	}
}

func TestGameRemovePlayer(t *testing.T) {
	players := make(map[string]*Player)
	players["john"] = &Player{Username: "john"}
	p := &Game{
		Players: players,
	}
	err := p.RemovePlayer(&Player{Username: "john"})
	if err != nil {
		t.Errorf("Failed to remove player : %v", err)
	}
}
