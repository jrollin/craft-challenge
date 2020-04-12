package rest

import (
	"github.com/google/uuid"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"log"
	"net/http"
)

type AddGameHandler struct {
	log    *log.Logger
	storer port_in.StoreGame
}

func NewAddGameHandler(log *log.Logger, storer port_in.StoreGame) *AddGameHandler {
	return &AddGameHandler{
		log:    log,
		storer: storer,
	}
}

func (g *AddGameHandler) AddGame(rw http.ResponseWriter, r *http.Request) {

	g.log.Printf("[DEBUG] add new game %s", r.Method)

	// decode request with anonymous struct
	t := struct {
		Code string `json:"code"`
	}{}
	err := utils.FromJSON(&t, r.Body)
	if err != nil {
		g.log.Printf("[ERROR] error decoding request %s", err)
		http.Error(rw, "Error processing request", http.StatusUnprocessableEntity)
		return
	}

	// create command from request
	id := uuid.New()
	ag, err := port_in.NewAddGame(id, t.Code)
	if err != nil {
		g.log.Printf("[ERROR] error creating command %s", err)
		http.Error(rw, "Invalid data provided", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = g.storer.Store(ag)
	if err != nil {
		g.log.Println("[ERROR] Store game failed", err)
		http.Error(rw, "Error storing game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
