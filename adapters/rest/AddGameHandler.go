package rest

import (
	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type AddGameHandler struct {
	l *log.Logger
	s port_in.AddGame
}

func NewAddGameHandler(log *log.Logger, adder port_in.AddGame) *AddGameHandler {
	return &AddGameHandler{
		l: log,
		s: adder,
	}
}

func (gh *AddGameHandler) AddGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] add new game %s", r.Method)

	// decode request with anonymous struct
	t := struct {
		Code string `json:"code"`
	}{}
	err := utils.FromJSON(&t, r.Body)
	if err != nil {
		gh.l.Printf("[ERROR] error decoding request %s", err)
		http.Error(rw, "Error processing request", http.StatusUnprocessableEntity)
		return
	}

	// create command from request
	id := uuid.New()
	ag, err := port_in.NewAddGameCommand(id, t.Code)
	if err != nil {
		gh.l.Printf("[ERROR] error creating command %s", err)
		http.Error(rw, "Invalid data provided", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.s.AddGame(ag)
	if err != nil {
		gh.l.Printf("[ERROR] adding game failed %s", err)
		http.Error(rw, "Error adding game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
