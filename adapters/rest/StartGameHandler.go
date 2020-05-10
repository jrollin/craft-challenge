package rest

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/domain/port_in/command"
	"github.com/jrollin/craft-challenge/domain/port_in/query"
)

type StartGameHandler struct {
	l *log.Logger
	f query.FindGameByCode
	s command.StartGame
}

func NewStartGameHandler(log *log.Logger, starter command.StartGame, finder query.FindGameByCode) *StartGameHandler {
	return &StartGameHandler{
		l: log,
		f: finder,
		s: starter,
	}
}

// swagger:route POST /games/{id}/start game addGameId
//
// Start a game
//
// Responses:
// 	default: genericErrorResponse
//  201: noContentResponse
//  422: validationErrorResponse
func (gh *StartGameHandler) StartGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] start a game %s", r.Method)

	vars := mux.Vars(r)
	id := vars["id"]

	c, err := command.NewStartGameCommand(uuid.MustParse(id))
	if err != nil {
		gh.l.Printf("[ERROR] error command request %s", err)
		http.Error(rw, "Error processing command", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.s.StartGame(c)
	if err != nil {
		gh.l.Printf("[ERROR] starting game failed %s", err)
		http.Error(rw, "Error starting game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
