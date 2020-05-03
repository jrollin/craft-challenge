package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jrollin/craft-challenge/adapters/rest/utils"
	"github.com/jrollin/craft-challenge/application/port_in"
	"github.com/jrollin/craft-challenge/domain"
)

type StartGameHandler struct {
	l *log.Logger
	f port_in.FindGame
	s port_in.StartGame
}

func NewStartGameHandler(log *log.Logger, starter port_in.StartGame, finder port_in.FindGame) *StartGameHandler {
	return &StartGameHandler{
		l: log,
		f: finder,
		s: starter,
	}
}

// swagger:route POST /games/{code}/start game addGameId
//
// Add a new game
//
// Responses:
// 	default: genericErrorResponse
//  201: noContentResponse
//  422: validationErrorResponse
func (gh *StartGameHandler) StartGame(rw http.ResponseWriter, r *http.Request) {

	gh.l.Printf("[DEBUG] start a game %s", r.Method)

	vars := mux.Vars(r)
	code := vars["code"]

	g, err := gh.f.FindByCode(domain.GameCode(code))
	if err != nil {
		gh.l.Printf("[ERROR] Game not found  %s", err)
		http.Error(rw, "Error finding game", http.StatusBadRequest)
		return
	}

	// decode request with anonymous struct
	t := &StartGameRequest{}

	err = utils.FromJSON(t, r.Body)
	if err != nil {
		gh.l.Printf("[ERROR] error decoding request %s", err)
		http.Error(rw, "Error processing request", http.StatusUnprocessableEntity)
		return
	}

	// call usecase with valid command
	err = gh.s.StartGame(g)
	if err != nil {
		gh.l.Printf("[ERROR] starting game failed %s", err)
		http.Error(rw, "Error starting game", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

// An StartGameRequest model.
//
// swagger:parameters addGameId
type StartGameRequest struct {
	// The code to submit
}
